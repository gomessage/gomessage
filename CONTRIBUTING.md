# Contributing Guide

本文档用于约定 GoMessage 的开发流程与分支策略，便于多人协作和稳定发布。

## 为什么单独建这个文档

README 主要面向“使用者”，而贡献与协作规范主要面向“开发者”。
在开源社区中，这类内容通常放在 `CONTRIBUTING.md`，并置于仓库根目录。

## 分支模型

- `master`：生产稳定分支，只接收已验证可发布的代码
- `release`：发布候选分支，用于发版前冻结、回归测试、版本校验
- `dev`：日常开发集成分支，功能开发最终先合入这里

## 日常开发流程

1. 从 `dev` 拉取最新代码
2. 基于 `dev` 创建功能分支
3. 开发完成后提交到功能分支并发起合并请求到 `dev`
4. 准备发版时，从 `dev` 合并到 `release`
5. 测试通过后，从 `release` 合并到 `master` 并打版本标签
6. 发版完成后，将 `release` 回灌到 `dev`

## 分支命名建议

- 功能开发：`feature/<name>`
- 缺陷修复：`fix/<name>`
- 紧急修复：`hotfix/<name>`
- 重构优化：`refactor/<name>`

## 合并策略

- 不在 `master` 直接开发
- 不在 `release` 做无关需求开发
- 通过合并请求进行代码评审后再合并
- 合并前确保基础构建和关键校验通过

## 版本与标签

- 版本发布以 `master` 上的 Tag 为准，格式建议：`vX.Y.Z`
- `release` 分支用于准备发布，不作为最终版本事实来源

## 热修复流程

1. 从 `master` 拉取 `hotfix/<name>` 分支
2. 修复后先合入 `master` 并打补丁版本 Tag
3. 将同一修复回灌到 `dev`
4. 若当前存在进行中的 `release`，同样合入 `release`

## 常用命令示例

```bash
git switch dev
git pull
git switch -c feature/xxx
git add .
git commit -m "feat: xxx"
git push -u origin feature/xxx
```

```bash
git switch release
git merge --no-ff dev
```

```bash
git switch master
git merge --no-ff release
git tag -a vX.Y.Z -m "release vX.Y.Z"
git push origin master --tags
```
