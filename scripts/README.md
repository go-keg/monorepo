# Makefile 命令

## 生成k8s部署文件
```bash
make k8s.config.${NAMESPACE}
```

## 打包
```bash
make ${SERVICE}.build
```

## 构建镜像
```bash
make ${SERVICE}.image
```

## 推送镜像
```bash
make ${SERVICE}.publish
```

## 推送公共镜像
```bash
make common_image.publish
```

## 部署
```bash
make ${SERVICE}.deploy
```