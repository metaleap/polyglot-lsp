package glot

import (
	"os"
	"os/exec"
	"strings"
)

type GenLangCmd struct {
	Disabled bool
	Cmd      string
	Env      map[string]string
	PerFile  bool
}

func (it *GenLangCmd) ok() bool {
	return it != nil && !it.Disabled
}

func (it *GenLangCmd) exec(gen *Gen, repl map[string]string) {
	if repl == nil {
		repl = map[string]string{}
	}
	for old, new := range map[string]string{
		"{dir}":           PathAbs(gen.dirPathDst),
		"{trash_dir}":     trashDir,
		"{pkg_file_name}": gen.Main.Lang.PkgFile,
		"{pkg_file_path}": PathAbs(gen.dirPathDst + "/" + gen.Main.Lang.PkgFile),
		"{user_home_dir}": os.Getenv("HOME"),
	} {
		if repl[old] == "" {
			repl[old] = new
		}
	}

	parts := Without(strings.Split(it.Cmd, " "), "")
	for i := 0; i < len(parts); i++ {
		if parts[i] == "{files}" {
			parts = append(parts[:i], append(gen.tracked.filesGenerated.code, parts[i+1:]...)...)
			i--
		} else {
			for old, new := range repl {
				parts[i] = strings.ReplaceAll(parts[i], old, new)
			}
		}
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	for name, value := range it.Env {
		for old, new := range repl {
			value = strings.ReplaceAll(value, old, new)
		}
		it.Env[name] = value // just for the `full_command` print-out further below
		cmd.Env = append(cmd.Env, name+"="+value)
	}
	cmd.Env = append(os.Environ(), cmd.Env...)
	{
		full_command := strings.Join(parts, " ")
		for name, value := range it.Env {
			full_command = name + "=" + value + "   " + full_command
		}
		if idx := strings.LastIndex(full_command, "   "); idx > 0 {
			full_command = full_command[:idx] + "\n\t" + full_command[idx+2:]
		}
		println(">>>", full_command)
	}
	if output, err := cmd.CombinedOutput(); err != nil {
		panic(err.Error() + ":\n" + string(output))
	}
}
