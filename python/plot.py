import argparse
import os

import matplotlib.pyplot as plt
import numpy as np


def read_data(filename, num_objects):
    all_positions = [[] for _ in range(num_objects)]

    with open(filename, "r") as f:
        for line_index, line in enumerate(f):
            line = line.strip()
            if not line:
                continue

            data = line.split(",")
            position = [
                float(data[0].strip()),
                float(data[1].strip()),
                float(data[2].strip()),
            ]

            object_index = line_index % num_objects
            all_positions[object_index].append(position)

    all_positions = [np.array(positions) for positions in all_positions]
    return all_positions


def plot_trajectory(all_positions, labels, output_path, show_labels=True):
    if not all_positions:
        print("Error: No positions to plot.")
        return

    fig = plt.figure(figsize=(12, 8))
    ax = fig.add_subplot(111, projection="3d")

    colors = plt.cm.tab20.colors
    for i, positions in enumerate(all_positions):
        x, y, z = positions[:, 0], positions[:, 1], positions[:, 2]
        ax.plot(x, y, z, color=colors[i % len(colors)], alpha=0.8, linewidth=2)
        ax.scatter(
            x[-1],
            y[-1],
            z[-1],
            color=colors[i % len(colors)],
            s=100,
            label=labels[i] if show_labels else None,
        )

    ax.set_xlabel("X", labelpad=15)
    ax.set_ylabel("Y", labelpad=15)
    ax.set_zlabel("Z", labelpad=15)
    ax.tick_params(axis="both", which="major", labelsize=10)
    ax.grid(True, linestyle="--", alpha=0.6)
    ax.legend(loc="center left", bbox_to_anchor=(1.1, 0.5), fontsize=10)
    plt.axis("equal")
    plt.title("3D Trajectories", fontsize=16, pad=20)
    plt.savefig(output_path, bbox_inches="tight", dpi=300)
    plt.close(fig)


def plot_2d_projection(all_positions, labels, output_path, show_labels=True):
    if not all_positions:
        print("Error: No positions to plot.")
        return

    plt.figure(figsize=(8, 8))

    colors = plt.cm.tab20.colors
    for i, positions in enumerate(all_positions):
        x, y = positions[:, 0], positions[:, 1]
        plt.plot(x, y, color=colors[i % len(colors)], alpha=0.8, linewidth=2)
        plt.scatter(
            x[-1],
            y[-1],
            color=colors[i % len(colors)],
            s=100,
            label=labels[i] if show_labels else None,
        )

    plt.xlabel("X", fontsize=12, labelpad=10)
    plt.ylabel("Y", fontsize=12, labelpad=10)
    plt.title("2D Projections (X vs Y)", fontsize=14, pad=15)
    plt.axis("equal")
    plt.grid(True, linestyle="--", alpha=0.6)
    plt.legend(loc="center left", bbox_to_anchor=(1.1, 0.5), fontsize=10)
    plt.savefig(output_path, bbox_inches="tight", dpi=300)
    plt.close()


def main():
    parser = argparse.ArgumentParser(description="Visualize trajectory data")
    parser.add_argument(
        "--filename", type=str, required=True, help="Input data file path"
    )
    parser.add_argument("--N", type=int, help="Number of objects")
    parser.add_argument(
        "--labels",
        type=str,
        help="Comma-separated labels for the objects (optional). "
        "Example: 'Sun,Earth,Moon'",
    )
    parser.add_argument(
        "--show-labels",
        type=str,
        choices=["True", "False"],
        help="Show or hide the object label in the output plots.",
    )

    args = parser.parse_args()

    show_labels = args.show_labels == "True" if args.show_labels else True

    all_positions = read_data(args.filename, args.N)
    num_objects = len(all_positions)

    if args.labels:
        labels = [label.strip() for label in args.labels.split(",")]
        if len(labels) != num_objects:
            print(
                f"Error: Number of labels ({len(labels)}) does not match "
                f"the number of objects ({num_objects})."
            )
            return
    else:
        labels = [f"Object {i+1}" for i in range(num_objects)]

    trajectory_plot_path = f"{args.filename}_trajectory.png"
    projection_plot_path = f"{args.filename}_2d_projection.png"

    plot_trajectory(all_positions, labels, trajectory_plot_path, show_labels)
    plot_2d_projection(all_positions, labels, projection_plot_path, show_labels)


if __name__ == "__main__":
    main()
