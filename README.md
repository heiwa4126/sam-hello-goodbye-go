# sam-hello-goodbye

1個のAWS Lambdaで複数のapi gwを受けるサンプル。GoLang版

[AWS LambdaでPython+Flask環境を作成する](https://ryosh.github.io/posts/20210421/)
みたいに、api gwでなんでもかんでも受ける、にする手もありそう。

これも参考: [go言語のwebフレームワークechoとpythonのwebフレームワークflaskの速度を比較してみた](https://qiita.com/shibacow/items/320679971bfd5d834b80)も参考


# インストール

SAMなのでいつもどおりに
```sh
sam build   # or `make`
sam deploy --guided   # 2回め以降は --guided ぬきで
```

でもこのプロジェクトはデプロイしないでも
```sh
sam build
sam local start-api
# tmuxだったら画面分割して
curl http://localhost:3000/hello
curl http://localhost:3000/hello/test
curl http://localhost:3000/goodbye/test
```
で十分。


# TODO

他のトリガ(Cloudwatchのschedule)なども混ぜてみる。
