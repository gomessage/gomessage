# Contributing Guide

本文档用于约定 GoMessage 的开发流程与分支策略，便于多人协作和稳定发布。

## 为什么单独建这个文档

README 主要面向“使用者”，而贡献与协作规范主要面向“开发者”。
在开源社区中，这类内容通常放在 `CONTRIBUTING.md`，并置于仓库根目录。

## 分支模型

- `main`：生产稳定分支，只接收已验证可发布的代码
- `dev`：日常开发集成分支，功能开发最终先合入这里
- `feature/*`：从 `dev` 拉出的功能分支，用于单个需求开发
- `fix/*`：从 `dev` 拉出的缺陷修复分支，用于非紧急问题修复
- `hotfix/*`：从 `main` 拉出的紧急修复分支，用于线上问题快速止血

## 日常开发流程

1. 从 `dev` 拉取最新代码
2. 基于 `dev` 创建功能分支
3. 开发完成后提交到功能分支并发起合并请求到 `dev`
4. 需要发布时，从 `dev` 合并到 `main`
5. 在 `main` 上完成发布校验并打版本标签
6. 若发布前需要单独冻结，可临时创建 `release/<version>` 分支，但不长期保留

## 分支命名建议

- 功能开发：`feature/<name>`
- 缺陷修复：`fix/<name>`
- 紧急修复：`hotfix/<name>`
- 重构优化：`refactor/<name>`

## 合并策略

- 不在 `main` 直接开发
- 不在 `dev` 直接堆叠长期未拆分的大需求
- 通过合并请求进行代码评审后再合并
- 合并前确保基础构建和关键校验通过
- `release/<version>` 仅在确有发版冻结需求时临时创建，发布后及时删除

## 版本与标签

- 版本发布以 `main` 上的 Tag 为准，格式建议：`vX.Y.Z`
- `main` 是唯一稳定事实来源，`dev` 用于持续集成与开发
- 如使用 `release/<version>`，它只作为临时发版分支，不作为长期主线

## 热修复流程

1. 从 `main` 拉取 `hotfix/<name>` 分支
2. 修复后先合入 `main` 并打补丁版本 Tag
3. 将同一修复回灌到 `dev`
4. 若当前存在临时 `release/<version>` 分支，同样合入该分支

## 常用命令示例

```bash
git switch dev
git pull
git switch -c feature/xxx
git add .
git commit -m "feat: xxx"
git push -u github feature/xxx
```

```bash
git switch main
git merge --no-ff dev
```

```bash
git tag -a vX.Y.Z -m "release vX.Y.Z"
git push github main --tags
```

```bash
git switch -c hotfix/xxx main
git add .
git commit -m "fix: xxx"
git push -u github hotfix/xxx
```
