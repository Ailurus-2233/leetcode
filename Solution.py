import unittest

class Solution:
    def add(self, a, b):
        return a + b


class SolutionTest(unittest.TestCase):

    solution = Solution()

    def test_add(self):
        self.assertEqual(self.solution.add(2, 3), 5)
        self.assertEqual(self.solution.add(-1, 1), 1)
        self.assertEqual(self.solution.add(0, 0), 1)

if __name__ == '__main__':
    unittest.main()