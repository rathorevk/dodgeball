# Dodgeball Game Solution
A small game simulation in Go

## Problem Description
Kids are playing a dodgeball game on a coordinate system where:
- Players are positioned at coordinates where x,y are 0, 10, or multiples of 10
- When a player receives the ball, they can pass it in compass directions (multiples of 45°)
- If multiple players are in the chosen direction, the ball goes to the nearest one
- The game continues until no valid pass can be made

## Input Format
```
number_of_test_cases
number_of_players
x1 y1
x2 y2
...
xn yn
starting_direction
starting_player (1-based)
```

## Output Format
For each test case:
```
number_of_throws ending_player
```
- `number_of_throws`: Number of throws excluding the starting throw
- `ending_player`: Player who has the ball when the game ends (1-based)

## Compass Directions
- N (North): 0°
- NE (Northeast): 45°
- E (East): 90°
- SE (Southeast): 135°
- S (South): 180°
- SW (Southwest): 225°
- W (West): 270°
- NW (Northwest): 315°

## How to Run

1. **Clone the Repository**  
    ```bash
    git clone <repository-url>
    cd dodgeball
    ```

2. **Build and Run**  
    If using Go:
    ```bash
    go run main.go < input.txt
    ```
    Replace `main.go` with your entry file and `input.txt` with your input file.

3. **Input Format**  
    Prepare your input file as described below.

4. **Example Command**  
    ```bash
    go run main.go < example_input.txt
    ```

## Algorithm
1. **Direction Detection**: Uses vector mathematics to determine if a player lies on a specific compass ray
2. **Distance Calculation**: Euclidean distance for finding the nearest player
3. **Game Simulation**: Continues passing until no valid target exists
4. **Cycle Detection**: Prevents infinite loops by tracking visited states

## Example
```
Input:
1
5
0 0
10 0
20 10
10 10
0 10
E
1

Output:
1 2
```

This means:
- Starting from player 1 at (0,0) going East
- Ball goes to player 2 at (10,0) - that's 1 throw
- Player 2 cannot pass further East (no players), so game ends
- Final result: 1 throw, ending at player 2

## Key Features
- Handles all 8 compass directions correctly
- Finds nearest player when multiple options exist
- Detects game end conditions
- Prevents infinite loops with cycle detection