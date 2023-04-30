### Android UI

#### UI布局

##### 线性布局LinearLayout

| 属性                        | 可选值                                                                                       | 说明               |
| ------------------------- | ----------------------------------------------------------------------------------------- | ---------------- |
| orientation               | 1. vertical:垂直排列，2.horizontal:水平排列                                                        |                  |
| layout_width,layout_hight | 1.match_parent:填充父容器的控件,2.wrap_content:跟据子视图宽高适应宽高3.指定固定高宽度                               | 必须指定             |
| gravity                   | 1.center:所有子视图相对于父容器居中，2.horizontal_center:子容器横向方向上相对于父容器居中。3.vertical_center:所有子视图纵向居中显示 | 决定子控件相对服务器的位置    |
| layout_gravity            | 上同                                                                                        | 决定该容器相对于它的父容器的位置 |
| weiht                     |                                                                                           | 按比例分配宽度和高度       |

##### 相对布局RelatevelLayout

相对于父元素 摆放控件的位置    

| 属性                       | 可选值        | 说明         |
| ------------------------ | ---------- | ---------- |
| layout_alignPatentTop    | ture/false | 相对于父容器顶部对齐 |
| layout_alignParentBottom | true/false |            |
| layout_alignParentLeft   | true/false |            |
| layout_alignParentRight  | true/false |            |
| layout_alignHorizontal   | true/false | 水平居中       |
| layout_alignVertical     | true/false | 垂直居中       |
| centerParent             | true/false | 居中         |

相对于兄弟元素摆放自己的位置 

| 属性               | 可选值        | 说明         |
| ---------------- | ---------- | ---------- |
| layout_ablove    | 组件id,@+id/ | 指定在其他控件的上方 |
| layout_below     |            |            |
| layout_toRightOf |            |            |
| layout_toLeftOf  |            |            |

相对于兄弟元素对齐方式

| 属性                 | 可选值   | 说明                 |
| ------------------ | ----- | ------------------ |
| layout_alignLeft   | @+id/ | 该控件的左边沿与指定控件左边对齐   |
| layout_alignRight  |       |                    |
| layout_alignTop    |       |                    |
| layout_alignBottom |       | 该控件的下边沿与指定控件的下边沿对齐 |

##### 帧布局FrameLayout

组件的默认位置都在左上角，组件可以重叠

| 属性                  | 可选值                                      | 说明         |
| ------------------- | ---------------------------------------- | ---------- |
| layout_gravity      | center/center_vertical/center_horizontal | 组件相对于父原件位置 |
| layout_marginLeft   | 具体数值dp                                   | 左侧外间距      |
| layout_marginRight  |                                          |            |
| layout_marginTop    |                                          |            |
| layout_marginBottom |                                          |            |

#### 控件

##### MaterialButton

谷歌SDK28推出的新控件，方便使用描边、圆角、前后置的图标等特殊按钮形状 

添加依赖

`implementation 'com.google.android.material:material:1.5.0'`

修改app主题（MaterialComponets下的都行）

`android:theme="@style/Theme.MaterialComponets.Light.NoAcationBar"`

去掉按钮的阴影，按钮更扁平化

`style="@style/Widget.MaterialComponents.Button.UnelevatedButton`

app出现闪退,解决方案修改app主题

`android:theme="@style/Theme.MaterialComponets.Light.NoAcationBar`

| 属性                 | 说明     | 参数                                           |
| ------------------ | ------ | -------------------------------------------- |
| backgroundTint     | 背景颜色   | 颜色                                           |
| backgroundTintMode | 着色模式   | add,multiply,screen,src_atop,src_in,src_over |
| strokeColor        | 描边颜色   |                                              |
| strokeWidth        | 描边宽度   |                                              |
| corenerRadius      | 圆角大小   |                                              |
| rippleColor        | 水波纹颜色  |                                              |
| icon               | 图标     |                                              |
| iconSize           | 图标大小   |                                              |
| iconGravity        | 着色重心   | start,end,textStart,textEnd                  |
| iconTint           | 图标着色   |                                              |
| iconTintMode       | 图标着色模式 | add,multiply,screen,src_atop,src_in,src_over |
| iconPadding        | 图标内部间距 |                                              |

##### MaterialButtonToggleGroup

类似一个linearLayout，可以添加多个MaterialButton构成单选或者多选按钮

| 属性                | 描述            | 参数         |
| ----------------- | ------------- | ---------- |
| checkedButton     | 是否默认选中        | true/false |
| singleSelecction  | 是否单项选择        |            |
| selectionRequired | 为true时，至少选择一个 |            |

监听按钮组

```kotlin
val toggleGroup = findViewById<MaterialButtonToggleGroup>(R.id.button_group)  
toggleGroup.addOnButtonCheckedListener { group, checkedId, isChecked ->  
    Log.e(  
        "MainActivity:$checkedId",  
        isChecked.toString()  
    )  
}
```

##### TextView

文本控件

常用属性:

| 属性                 | 说明           | 参数                     |
| ------------------ | ------------ | ---------------------- |
| text/textcolor     |              |                        |
| textappereance     | 文字外观         |                        |
| textcolorhighlight | 被选中文字的底色默认蓝色 |                        |
| textcolorlink      | 文字链接的颜色      |                        |
| textscalex         | 文字之间的间隔      | 默认1.0f                 |
| textsize           | 文字大小         | 15sp                   |
| textstyle          | 字形           | bold,italic,bolditalic |
| autolink           | 文字链接         | url,email,phone,map    |
| autotext           | 输入值拼写纠正      |                        |
| capitalize         | 设置英文首字母大写    |                        |
| cursorvisiable     | 设定光标显示/隐藏    | 默认显示                   |
| digits             | 设置允许输入的字符    | "123456789.+-*/%()"    |
| inputtype          | 设置文本的类型      |                        |
| password           | 以密码方式显示      |                        |
| typeface           | 设置字体         |                        |

##### ImageView

图片控件

| 属性         | 说明      | 示例                  |
| ---------- | ------- | ------------------- |
| src        | 显示图片的地址 | @drawable/icon_home |
| alpha      | 透明度     | 0-1                 |
| background | 背景颜色或图片 |                     |

scaleType：图片的缩放对控件的适配

-  center：显示在控件的中心，超过部分进行裁切处理。

- centerCrop：填满整个控件，等比例放到直到填满控件

- centerInside：显示完全整个图，通过按比例缩小

- fitCenter：把原图按比例扩大缩小到控件高度，居中显示

- fitEnd：把原图按比例扩大缩小到控件高度，显示下部分

- fitStart：把原图按比例扩大缩小到控件高度，显示上部分

- fitXY：把原图不按比例拉伸显示图片，填满控件显示

- matrix：不改变原图大小，从左上角开始绘制，超过部分进行裁切



##### RecycleView

它具有出色的性能(四级回收和复用机制)和拔插式的架构，大部分产品都有使用

###### LinerLayoutManager

列表布局,常用的纵向列表和横向列表
