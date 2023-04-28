# JavaIO流

## 字节流和处理流(包装流)

节点流和处理流的区别和联系

1. 节点流是底层流/低级流,直接跟数据源相接
   处理流(包装流)包装节点流，既可以消除不同节点流的实现差异，也可以提供更方便的方

2. 法来完成输入输出。[源码理解]处理流(也叫包装流)对节点流进行包装，使用了修饰器设

3. 计模式，不会直接与数据源相连 [模拟修饰器设计模式]

处理流的功能主要体现在以下两个方面:

1. 性能的提高:主要以增加缓冲的方式来提高输入输出的效率

2. 操作的便捷: 处理流可能提供了一系列便捷的方法来一次输入输出大批量的数据，使用更加灵活方便

![javaIO流文件操作实例详解图片示例 | 红颜丽人](https://ts1.cn.mm.bing.net/th/id/R-C.62329de02346cda22e6b3a01e03e99c7?rik=pugHWARKdreSSA&riu=http%3a%2f%2fimg.blog.csdn.net%2f20130826164223890&ehk=q0247nhYlS%2fIiLAQlhoA3c0Qs4jzuIJSO2C0o2LSnJY%3d&risl=&pid=ImgRaw&r=0)

## 文件

#### 一 、文件流

文件在程序中是以流的形式来操作的。

                               输入流：读取文件

   java程序(内存) <= 文件(磁盘)

                                 =>

                               输出流：写入文件

#### 二、文件常用操作

创建文件

`new File(String "路径")`

`new File(String | File "父路径","子路径")`

`file.createNewFile()`

获取文件信息

getName,getAbsolutePath,getParent,length,exists,isFile,isDirectory,delete

创建文件夹

`file.mkdir();`
`file.mkdirs();`

删除文件

`file.delete();`如果是文件夹，需要保证文件夹为空

## IO流原理及其分类

#### 一、java IO流原理

1. I/O是Input/Output的缩写，I/O技术是非常实用的技术，用于处理数据传输如读/写文件，网络通讯等。

2. Java程序中，对于数据的输入/输出操作以”流(stream)” 的方式进行。

3. java.io包下提供了各种“流”类和接口，用以获取不同种类的数据，并通过方法输入或输出数据

4. 输入input: 读取外部数据(磁盘、光盘等存储设备的数据) 到程序 (内存)中

5. 输出output: 将程序 (内存)数据输出到磁盘、光盘等存储设备中

#### 二、流的分类

1. 按操作数据单位不同分为: 字节流(8 bit)，字符流(按字符)

2. 按数据流的流向不同分为: 输入流，输出流

3. 按流的角色的不同分为: 节点流，处理流/包装流

| 抽象基类 | 字节流          | 字符流    |
| ---- | ------------ | ------ |
| 输入流  | InputStream  | Reader |
| 输出流  | OutputStream | Writer |

InputStream 常用的子类

- FilelnputStream: 文件输入流  

```java
FileInputStream fileInputStream = new FileInputStream(file)
//读取一个字节数据，如果返回-1则表示读取到文件末尾
while ((read = fileInputStream.read())!=-1){
//汉字由多个字节组成，单个字节转换char会变成乱码
       System.out.println((char) read);
      }
byte buffer[] =new byte[1024];
//一次读取1024个字节,此方法会造成阻塞,返回实际读取的字节数,读取完成返回-1
read = fileInputStream.read(buffer);
//将读取字节数组转换成字符串
String content = new String(buffer,0,read);
fileInputStream.close();
```

- FileReader

```java
InputStreamReader fileReader =new FileReader(file);
FileReader fileReader1 =new FileReader(file, Charset.defaultCharset());
fileReader.read();//读取到文件末尾返回-1
fileReader.read(new char[5],0,5);//读取5个字符
fileReader.close();
```

- FileOutputStream: 文件输出流

```java
//文件不存在会自动创建该文件,append:追加模式，默认是覆盖模式
fileOutputStream = new FileOutputStream(file,true);
byte buffer[] =new byte[1024];
//支持字符和整形
fileOutputStream.write('a');
String content = "fileOutputStream";
//将字符串转换成字节流写入到文件
fileOutputStream.write(content.getBytes());
//从字节数组写入len个长度字节,参数:off偏移量
fileOutputStream.write(buffer,0,3);
fileOutputStream.close();
```

- FileWriter(<mark>使用后必须close或者flush，否则写入找不到文件</mark>)

```java
OutputStreamWriter f = new FileWriter(fileName,true);//追加模式或者覆盖模式
FileWriter fileWriter =new FileWriter(new File(fileName));
char[] a ="是五个字符".toCharArray();
fileWriter.write('a');//写入一个字符
fileWriter.write(a,0,5);//写入5个字符
fileWriter.close();
```

- BufferedInputStream: 缓冲字节输入流

- ObjectInputStream: 对象字节输入流

#### 三、FileStream应用

文件拷贝(注:操作C盘需要管理员权限)

```java
fileInputStream = new FileInputStream(fileName);//要复制的文件
fileOutputStream =new FileOutputStream("124.txt");//新的文件名称
int readLength;
byte[] buffer =new byte[1024];
while ((readLength = fileInputStream.read(buffer)) != -1){
fileOutputStream.write(buffer,0,readLength);
}
fileInputStream.close();
fileOutputStream.close();
```

## 对象处理流

- ObjectInputStream：将二进制的数据读取为内存中的对象，也称为反序列化的过程。

- ObjectOutputStream：将对象或基本数据类型转化成二进制数据(包含值和数据类型)，即序列化的过程。

- 序列化 的好处在于可将任何实现了 Serializable 接口的对象转化为 字节数据 ， 使其在保存和传输时可被还原。

<mark>注：不能序列化 static 和 transient 修饰的成员变量</mark>,如果需要让某个对象支持序列化机制，则必须让对象所属的类及其属性是可序列化的，该类必须实现如下两个接口之一。否则，会抛出NotSerializableException异常： Serialization接口（常用）和 Externalizable


凡是实现Serializable接口的类都有一个表示序列化版本标识符的静态变量： private static final long serialVersionUID;
serialVersionUID的作用是，如果你把这个对象序列化之后，再反序列化时，会根据这个编号来识别原来第哪个person对象从而还原。

```java
FileOutputStream outputStream =new FileOutputStream("object.txt");  
objectOutputStream =new ObjectOutputStream(outputStream);  
objectOutputStream.writeObject(student);  
FileInputStream inStream = new FileInputStream("object.txt");  
p = new ObjectInputStream(inStream);  
Student m = (Student)p.readObject();
```
