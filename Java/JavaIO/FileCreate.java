import java.io.*;
import java.nio.charset.Charset;

public class FileCreate {

    public static void main(String[] args) {
        UserFile userFile = new UserFile();
        userFile.fileWriter();
    }

}

class UserFile {
    final String fileName = "abc.txt";

    public void createFile01() {
        String filePath = "E:\\test\\new.txt";
        File file = new File(filePath);
        try {
            file.createNewFile();
            file.delete();
            file.mkdir();
            file.mkdirs();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public void fileStreamOperate() {
        File file = new File(fileName);
        FileInputStream fileInputStream = null;
        try {
            if (!file.exists()) {
                file.createNewFile();
            }
            fileInputStream = new FileInputStream(file);
            int read;
//          read = fileInputStream.read();//读取一个字节数据，如果返回-1则表示读取到文件末尾
//            while ((read = fileInputStream.read())!=-1){
//                System.out.println((char) read);//汉字由多个字节组成，单个字节转换char会变成乱码
//            }
            byte buffer[] = new byte[1024];
            read = fileInputStream.read(buffer);//一次读取1024个字节,此方法会造成阻塞,返回实际读取的字节数，读取完成返回-1
            String content = new String(buffer, 0, read);//将读取字节数组转换成字符串
            fileInputStream.close();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public void fileReader() {
        File file = new File(fileName);
        try {
            InputStreamReader fileReader = new FileReader(file);
            FileReader fileReader1 = new FileReader(file, Charset.defaultCharset());
            fileReader.read();
            fileReader.read(new char[5], 0, 5);//读取5个字符
        } catch (FileNotFoundException e) {
            throw new RuntimeException(e);
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public void fileOutputStreamOperate() {
        File file = new File(fileName);
        FileOutputStream fileOutputStream = null;
        try {
            //文件不存在会自动创建该文件,append:追加模式，默认是覆盖模式
            fileOutputStream = new FileOutputStream(file, true);
            byte buffer[] = new byte[1024];
            //支持字符和整形
            fileOutputStream.write('a');
            String content = "fileOutputStream";
            //将字符串转换成字节流写入到文件
            fileOutputStream.write(content.getBytes());
            //从字节数组写入len个长度字节,参数:off偏移量
            fileOutputStream.write(buffer, 0, 3);
            fileOutputStream.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public void fileWriter() {
        try {
            OutputStreamWriter f = new FileWriter(fileName, true);
            FileWriter fileWriter = new FileWriter(fileName, false);
            char[] a = "是五个字符".toCharArray();
            fileWriter.write('a');//写入一个字符
            fileWriter.write(a, 0, 5);//写入5个字符
            fileWriter.flush();
            fileWriter.write("sfhahfafskajfksakfjsafsa");
            fileWriter.close();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    public void fileCopy() {
        FileOutputStream fileOutputStream = null;
        FileInputStream fileInputStream = null;
        try {
            fileInputStream = new FileInputStream(fileName);//要复制的文件
            fileOutputStream = new FileOutputStream("124.txt");//新的文件名称
            int readLength;
            byte[] buffer = new byte[1024];
            while ((readLength = fileInputStream.read(buffer)) != -1) {
                fileOutputStream.write(buffer, 0, readLength);
            }
            fileInputStream.close();
            fileOutputStream.close();
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }
}
