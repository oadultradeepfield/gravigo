# Gravigo

Gravigo is a Go program for simulating the orbit of N objects under their mutual gravitation. It is a porting of my two continuous previous projects written in C++. Visit [this repository](https://github.com/oadultradeepfield/n-body-orbit-simulation) for more details. Gravigo is meant to be the exact replica but with native support for orbit plotting without relying on Python's Matplotlib. It also supports concurrent simulation using goroutines to allow faster processing.

## Development Roadmap

- [x] Port the previous C++ implementation to Go with additional project structure improvement.
- [ ] Migrate the Python matplotlib visualization to [Gonum Plot](https://github.com/gonum/plot) for native experiences.
- [ ] Add concurrency support with goroutines in Runge-Kutta numerical integration for faster processing.

## Getting Started

1. Begin by cloning the repository and navigate to the respective directory

   ```bash
   git clone https://github.com/oadultradeepfield/gravigo
   cd gravigo
   ```

2. Download all dependencies using the following command:

   ```bash
   go mod tidy
   ```

3. Make sure you have Go installed. This project is built using Go version 1.23.4. You can run the simulation using the following command line:

   ```bash
   go run cmd/simulation/main.go config.json
   ```

   Please refer to the below section to set up the configuration file `config.json`,

## Usage

1.  From the previous section, you have seen how to run the simulation using the command line. Note that the configuration file is mandatory as it contains the data of the objects you would like to examine. Below is the example of how you can set up the file:

    ```json
    {
      "simulator_config": {
        "gravitational_constant": 6.6743e-11,
        "dt": 1000,
        "total_time": 3.16e7,
        "output_file": "example_sun_earth_lagrangian_points.txt"
      },
      "coordinate_type": "spherical",
      "bodies": [
        {
          "_name": "Sun",
          "mass": 1.989e30,
          "position": [0.0, 0.0, 1.5707963268],
          "velocity": [0.0, 0.0, 0.0]
        },
        {
          "_name": "Earth",
          "mass": 5.972e24,
          "position": [1.496e11, 0.0, 1.5707963268],
          "velocity": [0.0, 2.9788e4, 0.0]
        },
        {
          "_name": "L1",
          "mass": 6500,
          "position": [1.481e11, 0.0, 1.5707963268],
          "velocity": [0.0, 2.9489e4, 0.0]
        },
        {
          "_name": "L2",
          "mass": 6500,
          "position": [1.511e11, 0.0, 1.5707963268],
          "velocity": [0.0, 3.0087e4, 0.0]
        },
        {
          "_name": "L3",
          "mass": 6500,
          "position": [1.496e11, 3.1415926536, 1.5707963268],
          "velocity": [0.0, 2.978e4, 0.0]
        },
        {
          "_name": "L4",
          "mass": 6500,
          "position": [1.496e11, 1.0471975512, 1.5707963268],
          "velocity": [0.0, 2.978e4, 0.0]
        },
        {
          "_name": "L5",
          "mass": 6500,
          "position": [1.496e11, -1.0471975512, 1.5707963268],
          "velocity": [0.0, 2.978e4, 0.0]
        }
      ]
    }
    ```

    The parameter `dt` is the time step size used to estimate the integration. The smaller ones would generally yield more accurate results but require longer to process. In addition, using large values will sometimes cause incorrect orbit shapes since the error is high. If you use Cartesian coordinates, the three components of `position` and `velocity` should be simply X, Y, and Z. For spherical coordinates, they are R, Theta, and Phi. In this project, Theta is measured from the X-axis, while Phi is measured from the Z-axis down to the XY plane.

2.  As you have previously seen, you can run the simulation using the following command line:

    ```bash
    go run cmd/simulation/main.go config.json
    ```

    The output from this command is a `.txt` file located at the specified path in `config.json`:

    ```
    0.057882, 0.000005, -0.000000
    149599980722.057831, 89363998.228348, -0.763470
    148099980905.314911, 88466998.245267, -0.755815
    151099980527.170044, 90260998.210362, -0.771125
    -149599980722.042450, -89339999.755764, -0.763470
    74722619652.576691, 129602053710.282852, -0.763470
    74877361068.657425, -129512713712.054565, -0.763470
    ...
    ```

    Each line corresponds to the Cartesian coordinates of each object in the same order as the input. There are seven objects in the above example, so the line will repeat for the same objects every seven lines at the new time step.

3.  To be updated after integrating Gonum Plot.

## License

This project is licensed under the MIT License. See the [`LICENSE`](/LICENSE) file for details.
