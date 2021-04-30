import java.io.File;
import java.io.FileWriter;
import java.io.IOException;

public class GenerateGiantPhoneticMapFunction {
    public static void main(String[] args) {
        FileWriter writer = null;
        try {
            new File("GiantPhoneticMap.txt");
            writer = new FileWriter("GiantPhoneticMap.txt");
            writer.write("func GiantPhoneticMap() (res map[string]string) {\n");
            writer.write("    giantPhoneticMap := make(map[string]string)\n");
            for (int i = 0; i < 1000; i++) {
                writer.write("    giantPhoneticMap[\"");
                writer.write(String.valueOf(i));
                writer.write("\"] = \"");
                writer.write(GenerateTest.phonetic(i));
                writer.write("\"\n");
            }
            writer.write("	return giantPhoneticMap\n");
            writer.write("}");
        } catch (IOException e) {
            System.out.println(e);
        } finally {
            try {
                writer.close();
            } catch (IOException e) {
                System.out.println(e);
            }
        }
    }
}
