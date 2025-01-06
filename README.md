プロジェクトをクローン後に以下を実行してください

```
% make ini
```

各レイヤーの依存の方向としては

domain(entity->repository) -> usecase -> presentation

domain -> infrastructure
