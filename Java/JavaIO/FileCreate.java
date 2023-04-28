import java.io.*;
import java.nio.charset.Charset;
import java.time.LocalDate;
import java.time.Month;

public class FileCreate {

    public static void main(String[] args) {
        Student student =new Student();
        student.setGender('F');
        student.setName("可汗");
        student.setScore(66);
        LocalDate localDate = LocalDate.now();
        student.setBirthday(new Birthday(localDate.getYear(),localDate.getMonthValue(),localDate.getDayOfMonth()));
        ObjectOutputStream objectOutputStream = null;
        ObjectInputStream p=null;
        try{
FileOutputStream outputStream =new FileOutputStream("object.txt");
objectOutputStream =new ObjectOutputStream(outputStream);
objectOutputStream.writeObject(student);
FileInputStream inStream = new FileInputStream("object.txt");
p = new ObjectInputStream(inStream);
Student m = (Student)p.readObject();
            System.out.println(m);
        }catch (IOException e){
            e.printStackTrace();
        } catch (ClassNotFoundException e) {
            throw new RuntimeException(e);
        }finally {
            try {
                p.close();
                objectOutputStream.close();
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        }


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

class Student implements  Serializable{
    private String name;

    @Override
    public String toString() {
        return "Student{" +
                "name='" + name + '\'' +
                ", score=" + score +
                ", gender=" + gender +
                ", birthday=" + birthday +
                '}';
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getScore() {
        return score;
    }

    public void setScore(int score) {
        this.score = score;
    }

    public char getGender() {
        return gender;
    }

    public void setGender(char gender) {
        this.gender = gender;
    }

    public Birthday getBirthday() {
        return birthday;
    }

    public void setBirthday(Birthday birthday) {
        this.birthday = birthday;
    }

    private int score;
    private char gender;
    private Birthday birthday;


}
class Birthday implements Externalizable{
    private int year;

    public Birthday() {
        super();
    }

    @Override
    public String toString() {
        return "Birthday{" +
                "year=" + year +
                ", month=" + month +
                ", day=" + day +
                ", letterStates='" + letterStates + '\'' +
                ", num=" + num +
                '}';
    }

    public Birthday(int year, int month, int dayOfMonth) {
        this.year = year;
        this.month = month;
        this.day = dayOfMonth;
    }

    public int getYear() {
        return year;
    }

    public void setYear(int year) {
        this.year = year;
    }

    public int getMonth() {
        return month;
    }

    public void setMonth(int month) {
        this.month = month;
    }

    public int getDay() {
        return day;
    }

    public void setDay(int day) {
        this.day = day;
    }

    private int month;
    private int day;
    //用于定制地序列化此类
    private String letterStates ="123";
    private int num = 88;
    @Override
    public void writeExternal(ObjectOutput out) throws IOException {
        out.writeInt(year);
        out.writeInt(month);
        out.writeInt(day);
        out.writeObject("123");
        out.write(num); //在序列化的数据最后加个88
    }

    @Override
    public void readExternal(ObjectInput in) throws IOException, ClassNotFoundException {
        year = in.readInt();
        month = in.readInt();
        day = in.readInt();
        letterStates = (String)in.readObject();
        num = in.read();  //把数字88加进来
    }
}