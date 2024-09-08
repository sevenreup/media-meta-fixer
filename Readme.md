# Media Meta Fixer CLI

Media Meta Fixer is a command-line application written in Go that helps you organize your video files by renaming them based on their movie title metadata.

## Features

- Scans a specified folder for video files containing "tfpdl-" in their filename
- Uses MediaInfo to extract the movie title from each video file
- Renames the files using the extracted movie title
- Works across different operating systems (supports both "mediainfo" and "MediaInfo" commands)
- Provides informative output about the renaming process

## Prerequisites

Before you can use Media Meta Fixer, make sure you have the following installed:

1. Go (version 1.11 or later)
2. [MediaInfo CLI](https://mediaarea.net/en/MediaInfo) (must be accessible from the command line)

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/sevenreup/media-meta-fixer.git
   ```
2. Navigate to the project directory:
   ```
   cd media-meta-fixer
   ```
3. Build the application:
   ```
   go build -o media-meta-fixer
   ```

## Usage

Run the application by providing the path to the folder containing your video files:

```
./media-meta-fixer /path/to/your/video/folder
```

The application will scan the specified folder (and its subfolders) for video files with "tfpdl-" in their names. It will then attempt to rename these files based on their movie title metadata.

## Example

If you have a file named `tfpdl-awesome_movie_2021.mp4` and its movie title metadata is "The Awesome Movie", the file will be renamed to `The Awesome Movie.mp4`.

## Notes

- The application only processes files that have "tfpdl-" in their name.
- If MediaInfo can't extract a movie title or if the extraction fails, the file will not be renamed, and an error message will be displayed.
- Make sure you have the necessary permissions to rename files in the specified folder.

## Troubleshooting

If you encounter the error "MediaInfo command not found", ensure that MediaInfo is installed on your system and accessible from the command line. You can verify this by running `mediainfo --version` or `MediaInfo --version` in your terminal.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
