# 2017/06/06 のメモ

## メニュー

* 環境構築 - vim 編

## 環境構築 - vim 編

### オススメプラグイン

* [dein](https://github.com/Shougo/dein.vim)
  * vim のパッケージマネージャ
* [vim-go](https://github.com/fatih/vim-go)
  * go で開発を行うための諸々
* [neocomplete](https://github.com/Shougo/neocomplete.vim)
  * オートコンプリートの支援

### dein の導入

インストール手順の詳細は公式ページ参照。
以下、ちょっとつまづきそうなところを補足。

* インストール用のスクリプト

下記の `{specify the installation directory}` のところはお好みで良いが、
例えば `.cache/dein` 等と指定する。

```bash
$ curl https://raw.githubusercontent.com/Shougo/dein.vim/master/bin/installer.sh > installer.sh
$ sh ./installer.sh {specify the installation directory}
```

* .vimrc の設定は、当初はとりあえず公式ページの .vimrc の記載をまるまる用いる。

```vim
if &compatible
  set nocompatible
endif
set runtimepath+={各々の HOME へのパス}/.cache/dein/repos/github.com/Shougo/dein.vim

if dein#load_state('{各々の HOME へのパス}/.cache/dein')
  call dein#begin('{各々の HOME へのパス}/.cache/dein')

  call dein#add('Shougo/dein.vim')

# ここにインストールしたいプラグインを追記していく

  call dein#end()
  call dein#save_state()
endif

filetype plugin indent on
syntax enable

" install not installed plugins on startup.
if dein#check_install()
  call dein#install()
endif
```

### vim-go の導入

* dein を用いてインストールする。
プラグインを追記する箇所に、以下のように書く。

```vim
call dein#add('fatih/vim-go')
```

* インストール後、必要なバイナリをダウンロードする。
vim を開いて、`:GoInstallBinaries` と入力する。


### vim-go の機能いろいろ

覚えておくとちょっと便利な vim-go の機能。

* :GoMetaLinter

色んな種類の lint を一気にかけてくれる。遅いのでたまにやると良さそう。

* :GoDoc 

ドキュメントを参照する。標準ライブラリの関数の使い方が分からないときなどに。
見たい関数にカーソルをあてた状態で `:GoDoc` と入力する。

* ハイライト

各所に色が付いて見やすくなる。.vimrc に以下のような記載をする。

```vim
let g:go_highlight_functions         = 1
let g:go_highlight_methods           = 1
let g:go_highlight_structs           = 1
let g:go_highlight_operators         = 1
let g:go_highlight_build_constraints = 1
```

* 自動 gofmt

gofmt も自動でしてくれる。以下の設定を .vimrc に記載する。
gofmt にプラスアルファの機能がついている goimports を設定するのがオススメ。

```vim
let g:go_fmt_autosave = 1
let g:go_fmt_command  = "goimports"
```

* 補完機能

ctrl-x ctrl-o の順番押すと補完候補が表示される。
押しにくかったら特定のキーにバインドするのも良いかと。
例えば、以下のような設定を .vimrc に記載する。

```vim
imap <C-Space> <C-x><C-o>
```

### neocomplete の導入 (option)

特定の文字列を打ったら自動で補完が出るようになるプラグイン。
導入はちょっと手間なので、時間があったら。

* vim を +lua でビルドする。
* vimproc、neocomplete を dein を使ってインストールする。
* オムニ補完の設定 .vimrc に記載する。

```vim
let g:neocomplete#force_omni_input_patterns.go = '\h\w\.\w*'
```


