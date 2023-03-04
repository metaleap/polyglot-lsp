lspver=3.17
vscver=1.75.1


curdir=$PWD
owndir="prereq"
if [ $0 != "$owndir/_fetch.sh" ]
then
    echo "$0: run from \$repoDir/gen, not from $curdir"
    exit 1
fi


# fetch LSP stuff

wget -O $owndir/lsp_$lspver.json https://github.com/microsoft/language-server-protocol/raw/gh-pages/_specifications/lsp/$lspver/metaModel/metaModel.json
wget -O $owndir/lsp_$lspver.schema.json https://github.com/microsoft/language-server-protocol/raw/gh-pages/_specifications/lsp/$lspver/metaModel/metaModel.schema.json


# fetch VSCodium stuff

wget -O $owndir/vsc_$vscver.d.ts https://github.com/microsoft/vscode/raw/$vscver/src/vscode-dts/vscode.d.ts
