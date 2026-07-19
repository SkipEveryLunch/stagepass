# これはなに？

Go + Kubernetesによるマイクロサービス構成を習得するための学習用PJ

# 起動方法

## 1. イメージをビルド

ここで付けたタグをもとにk8sのマニフェスト内で参照できる
（Orbstackでビルドしたイメージは、Orbstack組み込みのk8sで参照できる）

```:bash
docker build -t stagepass:0.1.0 .
```

## 2. Kubernetesリソースの立ち上げ

```:bash
kubectl apply -f deploy/stagepass.yaml
```

ちなみに片付けは以下で行う

```:bash
kubectl delete -f deploy/stagepass.yaml
```

## 3. ローカル->kubernetesのポートフォワード

```:bash
kubectl port-forward service/stagepass 8080:80
```

metadata nameがstagepassのserviceへと転送する
