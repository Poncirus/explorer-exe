# explorer-exe

使用exe代替文件夹快捷方式，让windows磁贴可定制

## 下载

在 [release](https://github.com/LiaoHanwen/explorer-exe/releases) 中下载

## 使用

1. 解压文件
2. 修改 path.txt ，填入目标文件夹路径
3. 将 explorer-exe.exe 和 explorer-exe.VisualElementsManifest.xml 中的 explorer-exe 修改为想要显示的名称
4. 配置 *.VisualElementsManifest.xml 文件

## VisualElementsManifest.xml 文件配置

``` xml
<Application xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
		<VisualElements
				BackgroundColor="#000000"           // 背景颜色
				ShowNameOnSquare150x150Logo="on"    // 中号图标是否显示名称 "on" | "off"
				Square150x150Logo="template.png"    // 中号图标
				Square70x70Logo="template.png"      // 小号图标
				ForegroundText="light" />           // 字体颜色 "light" | "dart"
</Application>
```
