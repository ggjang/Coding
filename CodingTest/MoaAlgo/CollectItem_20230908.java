package CodingTest.MoaAlgo;

import java.util.Arrays;

/**
 * CollectItem_20230908
 */
public class CollectItem_20230908 {
    public static void main(String[] args) {
        Solution solution = new Solution();
        int[][] rectangles = {{1,1,7,4},{3,2,5,5},{4,3,6,9},{2,6,8,8}};

        int result = solution.solution(rectangles,1,3,7,8);
        System.out.printf("result:%d", result);
    }
}



class Solution {
    public int solution(int[][] rectangle, int characterX, int characterY, int itemX, int itemY) {
        int answer = 0;
        getAllValidPoint(rectangle);
        // while (characterX != itemX || characterY != itemY) {
            
        // }
        return answer;
    }

    public int[][] getAllValidPoint(int[][] rectangle) {
        int[][][] allPoint = new int[rectangle.length][][];
        int all_point_count = 0;
        for (int[] point : rectangle) {
            allPoint[all_point_count++] = getAllPointRectangle(point);
        }
          System.out.println(Arrays.deepToString(allPoint));

        return allPoint[0];
    }

    public int[][] getAllPointRectangle(int[] rectangle) {
        int x_point_count = rectangle[2] - rectangle[0] + 1;
        int y_point_count = rectangle[3] - rectangle[1] + 1;
        int[][] result = new int[x_point_count * 2 + y_point_count * 2 - 4][2];
        int result_count = 0;
        System.out.printf("x count:%d, y count:%d, point cap:%d\n", x_point_count, y_point_count, result.length);
        for(int x = rectangle[0]; x <= rectangle[2]; x++) {
            result[result_count][0] = x;
            result[result_count++][1] = rectangle[1];
            result[result_count][0] = x;
            result[result_count++][1] = rectangle[3];
        }

        for(int y = rectangle[1] + 1; y <= rectangle[3] - 1; y++) {
            result[result_count][0] = rectangle[0];
            result[result_count++][1] = y;
            result[result_count][0] = rectangle[2];
            result[result_count++][1] = y;
        }

        return result;
    }

    // public int[] isCross(int[][] rectangle) {
        
    // }
}