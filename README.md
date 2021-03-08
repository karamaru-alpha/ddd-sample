# ddd-sample

DDD の実装サンプル


## domain

- Entity

ドメイン知識を持つ。識別子によって区別される可変のドメインオブジェクト。
ドメイン知識: ソフトウェアが存在しなくても、現実世界で行われる事。

- 値オブジェクト


Entityの構成要素として採用。constructorでバリデーション。
ロジックの散財・不正代入を防いだり、コードの表現力を増加させる効能がある。

- ドメインサービス


値オブジェクトやエンティティに実装すべきでないドメインの振る舞いを記述

- Factory


ライフサイクルにおける複雑な処理をカプセル化(UserEntity生成など)

- Repository

永続化のためのInterface

## infrastructure

永続化・再構築をDBと通信して実現する
Entityからgormタグを付与した構造体にDTOしている。

## application

ドメインオブジェクトを用いてusecaseを実現する進行役。

CleanArchtectureを意識して、application層の入力/出力値を定義したり、Interfaceにinteractorって名前をつけたりしている。
JWTトークンの作成やパスワードのHash化などの技術要件もこの層のサービスとして扱う。


## interfaces

リクエストの送受信・アプリケーションサービスの呼び出しを行う
