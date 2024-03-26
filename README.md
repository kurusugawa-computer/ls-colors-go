# kciguild-template-go
kciguild の go プロジェクトテンプレートです。 
Go の拡張機能のインストールや詳しい説明は https://kurusugawa.jp/confluence/pages/viewpage.action?pageId=1135936257 を参照してください。

## やるべきこと
### 1. [快適な開発のための VScode 設定方法](https://kurusugawa.jp/confluence/pages/viewpage.action?pageId=1135936257) を一読する

### 2. コンテナ内で VScode を開く
　まずは適当なディレクトリを VScode で開いてください。次に左下の水色のボタンをクリックし、表示されたダイアログの中から `Reopen in Container` を選択してください。これにより、ローカルで開いていたディレクトリをコンテナ上で開くことが可能となります。以降はコンテナ内の操作となります。

### 3. 周辺ツールをインストールする
　https://kurusugawa.jp/confluence/pages/viewpage.action?pageId=1135936257 を参考にしてください。

### 4. gokci をインストールする
　https://github.com/kurusugawa-computer/kciguild-gokci の説明に従って `gokci` をインストールしてください。

## 使い方 (自動)

[kurusugawa-computer/kciguild-init](https://github.com/kurusugawa-computer/kciguild-init)が使えます

```console
$ go get -u github.com/kurusugawa-computer/kciguild-init/cmd/kciguild-init
$ kciguild-init -description "This is test repository" foo-var
Reading ~/.config/hub
Creating repository: kurusugawa-computer/kciguild-foo-var
Cloning kciguild-foo-var into in-memory file system
Rewriting README.md and creating commit
Pushing changes
Adding branch protection rules for master
Adding branch protection rules for develop
Granting permission "push" to kci-guild
Granting permission "admin" to kci-shine

	URL: https://github.com/kurusugawa-computer/kciguild-foo-var

```

## 使い方 (手動)

https://github.com/kurusugawa-computer/kciguild-template の、「Use this template」をクリック

以下の設定を行う

### 手動で設定してください

#### Settings > Collaborators& teams

Add a team より、

 * kci-shine に Admin
 * kci-guild に Write

を付与

#### Settings > Branches > Branch protection rules

1． Add rules をクリック

2. Branch name pattern に master と入力

3. 以下項目にチェックする

 * Require pull request reviews before merging
 * Dismiss stale pull request approvals when new commits are pushed
 * Require review from Code Owners

4. Create をクリック

5. Add rules をクリック

6. Branch name pattern に develop と入力し、masterと同様にチェックを入れてCreate

#### README.mdの内容の削除

以上が完了したらこのREADME.mdの一番上の**kciguild-template**をリポジトリ名にし、他の部分を消してコミットしてください。