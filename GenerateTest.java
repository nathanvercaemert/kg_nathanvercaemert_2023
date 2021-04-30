import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.Random;

/**
 * Used to generate the files that will be used to
 * test the assessment at a large scale. This program
 * generates two text files: command.txt contains the
 * command that can be run to execute the Golang
 * program, and expected.txt is the expected output
 * of the program based on the arguments (that are
 * generated as part of command.txt.)
 */
public class GenerateTest {
    public static void main (String[] args) {
        FileWriter commandWriter = null;
        FileWriter expectedWriter = null;
        try {
            new File("command.txt");
            commandWriter = new FileWriter("command.txt");
            new File("expected.txt");
            expectedWriter = new FileWriter("expected.txt");
            commandWriter.write("go run main.go");
            Random rand = new Random();
            int iterations = 5000;
            for (int i = 0; i < iterations; i++) {
                commandWriter.write(" ");
                int nextNumber = rand.nextInt(1000);
                commandWriter.write(String.valueOf(nextNumber));
                expectedWriter.write(phonetic(nextNumber));
                if (i < iterations - 1) {
                    expectedWriter.write(", ");
                }
            }
            commandWriter.write(" > output.txt");
        } catch (IOException e) {
            System.out.println(e);
        } finally {
            try {
                commandWriter.close();
                expectedWriter.close();
            } catch (IOException e) {
                System.out.println(e);
            }
        }
    }

    public static String phonetic(int number) {
        if (number == 0) {
            return "Zero";
        }
        StringBuilder builder = new StringBuilder();
        StringBuilder builder2 = new StringBuilder();
        int sameNumber = number;
        while (sameNumber > 0) {
            builder2 = new StringBuilder();
            int modulo = sameNumber % 10;
            sameNumber /= 10;
            switch (modulo) {
                case 1: builder2.append("One"); break;
                case 2: builder2.append("Two"); break;
                case 3: builder2.append("Three"); break;
                case 4: builder2.append("Four"); break;
                case 5: builder2.append("Five"); break;
                case 6: builder2.append("Six"); break;
                case 7: builder2.append("Seven"); break;
                case 8: builder2.append("Eight"); break;
                case 9: builder2.append("Nine"); break;
                case 0: builder2.append("Zero"); break;
            }
            builder2.append(builder.toString());
            builder = new StringBuilder(builder2.toString());
        }
        return builder.toString();
    }
}