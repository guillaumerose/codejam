import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.assertj.core.util.Lists;
import org.jscience.mathematics.number.Real;

import static org.jscience.mathematics.number.Real.*;

public class ProblemC
{
    public static void main(String[] args)
    {
        Pattern pattern = Pattern.compile("([0-9]?[0-9]?[0-9])[.]");
        Real.setExactPrecision(500);
        List<Integer> ints = Lists.newArrayList(5, 12, 7, 4, 2, 22, 13, 27, 25, 6, 24, 30, 8, 3, 10, 18, 29, 23, 20, 26, 16, 21, 14, 19, 11, 17, 15, 28, 9);
        int i = 1;
        for (Integer pow : ints)
        {
            System.out.print("Case #" + i + ": ");
            Real value = valueOf(3).plus(valueOf(5).sqrt()).pow(pow);
            Matcher matcher = pattern.matcher(value.toText());
            matcher.find();
            String found = matcher.group(1);
            if (found.length() == 1)
            {
                System.out.print("00" + found);
            }
            else if (found.length() == 2)
            {
                System.out.print("0" + found);
            }
            else
            {
                System.out.print(found);
            }
            System.out.println();
            i++;
        }
    }
}
