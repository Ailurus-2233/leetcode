package q2;

import java.util.Arrays;
import java.util.Collection;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.Parameterized;

import junit.framework.TestCase;

@RunWith(Parameterized.class)
public class SolutionTest extends TestCase {

    // 相关参数
    // eg :
    // private int[] nums;
    // private int target;
    // private int[] expected;

    // 构造函数的输入参数，与上面的成员变量一一对应
    public SolutionTest() {

    }

    @Parameterized.Parameters
    public static Collection<Object[]> data() {
        return Arrays.asList(new Object[][]{
            // eg : {new int[]{2, 7, 11, 15}, 9, new int[]{0, 1}},
        });
    }

    private Solution solution = new Solution();
    @Test
    public void testSolution() {
        // eg : int[] result = solution.twoSum(nums, target);
        // assertEquals(expected, result);
    }
}