ver_lsp=3.17
ver_vsc=1.75.1


curdir=$PWD
owndir="prereq"
if [ $0 != "$owndir/_fetch.sh" ]
then
    echo "$0: run from \$repoDir/gen, not from $curdir"
    exit 1
fi


# fetch LSP stuff

wget -O $owndir/lsp_$ver_lsp.json https://github.com/microsoft/language-server-protocol/raw/gh-pages/_specifications/lsp/$ver_lsp/metaModel/metaModel.json
wget -O $owndir/lsp_$ver_lsp.schema.json https://github.com/microsoft/language-server-protocol/raw/gh-pages/_specifications/lsp/$ver_lsp/metaModel/metaModel.schema.json


# fetch VSCodium stuff

wget -O $owndir/vsc_$ver_vsc.d.ts https://github.com/microsoft/vscode/raw/$ver_vsc/src/vscode-dts/vscode.d.ts
