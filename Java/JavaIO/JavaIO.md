# JavaIO流

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
