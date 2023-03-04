package main //nolint

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	glot "polyglot-vsx-and-lsp/pkg"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestLoadMetaModelAndRoundTripCompareWithOrig(t *testing.T) {
	if err := os.Chdir("../.."); err != nil {
		t.Fatal(err)
	}
	for _, ver := range glot.Vers("lsp_", ".schema.json") {
		file_path := glot.SchemasDir + "/lsp_" + ver + ".json"
		meta_model := glot.LoadFromJSONFile[MetaModel](file_path)
		map_ours, map_orig := map[string]any{}, glot.LoadFromJSONFile[map[string]any](file_path)
		src, err := json.Marshal(meta_model)
		if err != nil {
			t.Fatal(err)
		}
		if err = json.Unmarshal(src, &map_ours); err != nil {
			t.Fatal(err)
		}
		for _, key := range []string{
			"enumerations",
			"structures",
			"typeAliases",
			"notifications",
			"requests",
		} {
			println(">>>>", key)
			for i, item := range map_orig[key].([]any) {
				m := item.(map[string]any)
				println(fmt.Sprintf(">>>>>>>>%v<<<<<<<<%v", m["name"], m["method"]))
				diff := cmp.Diff(map_orig[key].([]any)[i], map_ours[key].([]any)[i], cmpopts.EquateEmpty())
				if diff != "" {
					t.Fatal(diff)
					return
				}
			}
		}

	}
}
