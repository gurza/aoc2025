// Package day9 implements solutions for the Day 9 puzzle.
//
// # Part 1
//
// Part 1 scans all pairs of tiles in the original coordinate space.
// For each pair, it builds the axis-aligned rectangle defined by those two
// tiles, computes its area, and tracks the maximum area observed.
// This is a direct brute-force search over all candidate rectangles.
//
// # Part 2
//
// Part 2 uses a coordinate-compressed grid and several computational geometry
// steps to efficiently find the largest axis-aligned rectangle fully contained
// within the polygon formed by the input tiles.
//
// The algorithm consists of the following stages:
//
//  1. Coordinate Compression
//     All tile x/y coordinates are collected and mapped into dense integer
//     indices. This reduces the working domain from potentially large
//     coordinate ranges to a compact grid while preserving the geometric
//     order of coordinates.
//
//  2. Polygon Rasterization on the Compressed Grid
//     Each polygon vertex is converted into compressed coordinates. Edges
//     between consecutive vertices are drawn onto the grid by marking the
//     corresponding cells as inside. This produces a discrete representation
//     of the polygon boundary in the compressed space.
//
//  3. Flood-Fill Classification of Exterior Space
//     A breadth-first flood fill starts from the origin cell (0,0) in
//     compressed coordinates. Every reachable unknown grid cell is
//     reclassified as outside. This cleanly separates exterior cells from
//     interior and boundary cells enclosed by the polygon.
//
//  4. 2D Prefix-Sum Grid (Summed-Area Table)
//     The grid is transformed into a 2D prefix-sum table where each cell
//     stores the total count of non-outside cells in the rectangle from the
//     origin to that cell. This enables constant-time queries for how many
//     "inside-ish" cells lie within any axis-aligned rectangle on the
//     compressed grid.
//
//  5. Rectangle Containment Check and Area Extraction
//     For each pair of original tiles, the corresponding rectangle in the
//     compressed grid is checked using the prefix-sum table. If the rectangle
//     contains only non-outside cells, it is considered fully inside the
//     polygon. Its real geometric width and height are computed using the
//     original coordinates, and its area is compared against the current
//     maximum. The largest such area is returned as the Part 2 result.
package day9
