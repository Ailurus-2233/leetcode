namespace SolutionTest;

[TestClass]
public class UnitTest1
{
    Solution.Solution solution = new Solution.Solution();

    [TestMethod]
    [DynamicData(nameof(TestCase), DynamicDataSourceType.Method)]
    public void TestTwoSum(int[] nums, int target, int[] expected)
    {
        var result = solution.TwoSum(nums, target);
        Assert.IsTrue(AreEqual(result, expected));
    }

    private static IEnumerable<object[]> TestCase()
    {
        yield return new object[] { new int[] { 2, 7, 11, 15 }, 9, new int[] { 0, 1 } };
        yield return new object[] { new int[] { 3, 2, 4 }, 6, new int[] { 1, 2 } };
        yield return new object[] { new int[] { 3, 3 }, 6, new int[] { 0, 1 } };
    }

    private bool AreEqual(int[] nums1, int[] nums2)
    {
        if (nums1.Length != nums2.Length)
        {
            return false;
        }

        for (int i = 0; i < nums1.Length; i++)
        {
            if (nums1[i] != nums2[i])
            {
                return false;
            }
        }

        return true;
    }
}