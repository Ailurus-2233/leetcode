class Solution {
    public String replaceSpace(String s) {
        // return s.replace(" ", "%20");
        StringBuffer sb = new StringBuffer();
        for (char c: s.toCharArray()) {
            if (c == ' ') {
                sb.append("%20");
            } else {
                sb.append(c);
            }
        }
        return sb.toString();
    }
}