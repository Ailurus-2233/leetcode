import unittest
from solution import Solution


class SolutionTest(unittest.TestCase):
    solution = Solution()

    def test_add(self):
        self.assertEqual(self.solution.add(2, 3), 5)
        self.assertEqual(self.solution.add(-1, 1), 0)
        self.assertEqual(self.solution.add(0, 0), 0)


if __name__ == '__main__':
    unittest.main()
