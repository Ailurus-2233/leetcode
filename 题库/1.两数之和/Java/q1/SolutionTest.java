package q1;

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
    private int[] nums;
    private int target;
    private int[] expected;

    // 构造函数的输入参数，与上面的成员变量一一对应
    public SolutionTest(int[] nums, int target, int[] expected) {
        this.nums = nums;
        this.target = target;
        this.expected = expected;
    }

    @Parameterized.Parameters
    public static Collection<Object[]> data() {
        return Arrays.asList(new Object[][] {
                { new int[] { 2, 7, 11, 15 }, 9, new int[] { 0, 1 } },
                { new int[] { 3, 2, 4 }, 6, new int[] { 1, 2 } },
                { new int[] { 3, 3 }, 6, new int[] { 0, 1 } },
        });
    }

    private Solution solution = new Solution();

    @Test
    public void testSolution() {
        int[] result = solution.twoSum(nums, target);
        assertTrue(equals(result, expected));
    }

    private boolean equals(int[] expected, int[] result) {
        if (expected.length != result.length) {
            return false;
        }
        for (int i = 0; i < expected.length; i++) {
            if (expected[i] != result[i]) {
                return false;
            }
        }
        return true;
    }
}