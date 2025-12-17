# moon-phase
A beautiful GTK3 desktop application for displaying current moon phases with Wiccan magical correspondences.
![Moon Phase App](https://img.shields.io/badge/GTK-3.24-blue)
![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8)
![License](https://img.shields.io/badge/license-MIT-green)

## Features

- 🌑 **Accurate Moon Phase Calculation** - Uses Julian Day Number algorithm for precise astronomical calculations
- 🎨 **Beautiful Cairo Graphics** - Real-time rendering of moon phases with shadow effects
- ✨ **Wiccan Correspondences** - Each moon phase shows its magical meaning according to modern Wiccan tradition
- 🔄 **Live Updates** - Refresh button to update all information instantly
- 🖥️ **Native GTK3 Interface** - Clean, native Linux desktop integration
- 📊 **Detailed Information** - Shows illumination percentage, lunar age, current date/time

## Moon Phase Correspondences

| Phase | Wiccan Meaning |
|-------|----------------|
| 🌑 New Moon | ✨ New beginnings - Meditation and planning |
| 🌒 Waxing Crescent | ✨ Growth - Attraction and prosperity magic |
| 🌓 First Quarter | ✨ Action - Overcoming obstacles and decision-making |
| 🌔 Waxing Gibbous | ✨ Refinement - Preparation and fine-tuning |
| 🌕 Full Moon | ✨ Culmination - Maximum magical power |
| 🌖 Waning Gibbous | ✨ Gratitude - Sharing and thanksgiving |
| 🌗 Last Quarter | ✨ Release - Banishment and liberation |
| 🌘 Waning Crescent | ✨ Purification - Closing cycles |

## Screenshots

The application displays:
- Current moon phase with emoji representation (🌑🌒🌓🌔🌕🌖🌗🌘)
- Graphical moon rendering with accurate shadow positioning
- Phase name in Italian
- Illumination percentage
- Lunar age in days
- Current date and time
- Wiccan magical correspondence in italic

## Installation

### Prerequisites

**Arch Linux:**
```bash
sudo pacman -S go gtk3 clang
```

**Ubuntu/Debian:**
```bash
sudo apt install golang libgtk-3-dev clang
```

**Fedora:**
```bash
sudo dnf install golang gtk3-devel clang
```

### Dependencies

- Go 1.21+ (tested with Go 1.25.5)
- GTK3 3.24+
- Cairo 1.18+
- Clang compiler (required - see Compilation Notes)
- gotk3 v0.6.1 (automatically installed via `go build`)

## Building from Source

⚠️ **IMPORTANT: You must use Clang to compile this project.**

GCC 15.2.1 has a critical bug that causes segmentation faults when compiling GTK3 CGO bindings. Use Clang instead:

```bash
# Clone the repository
git clone https://github.com/yourusername/moon-phase.git
cd moon-phase

# Build with Clang (REQUIRED)
CC=clang CGO_ENABLED=1 go build -o moon-phase

# Or use the provided build script
./build.sh
```

### Quick Build Script

The project includes `build.sh` for convenience:

```bash
chmod +x build.sh
./build.sh
```

## Usage

### Running the Application

```bash
./moon-phase
```

### Desktop Integration

A desktop launcher file is included for XFCE4 panel integration. To install:

```bash
# Copy desktop file
cp moon-phase.desktop ~/.local/share/applications/

# Update desktop database
update-desktop-database ~/.local/share/applications/
```

The application will appear in your application menu under "Utility" or "Science" categories.

## Technical Details

### Moon Phase Calculation

The application uses precise astronomical calculations based on:
- **Synodic Month**: 29.530588853 days
- **Reference New Moon**: January 6, 2000 (Julian Day 2451549.5)
- **Julian Day Number Algorithm**: Accurate date-to-JD conversion

### Phase Thresholds

Astronomical thresholds for phase determination:
- New Moon: < 1.84566 days
- Waxing Crescent: < 5.53 days
- First Quarter: < 7.38264 days
- Waxing Gibbous: < 12.91 days
- Full Moon: < 14.76529 days
- Waning Gibbous: < 20.30 days
- Last Quarter: < 22.14794 days
- Waning Crescent: < 27.69 days

### Graphics Rendering

- Uses Cairo for 2D graphics rendering
- Real-time shadow calculation based on lunar age
- Smooth gradient rendering for realistic moon appearance
- Clipping masks for accurate shadow positioning

## Known Issues

### GCC Compilation Failure

**Problem:**
```
gcc: internal compiler error: Segmentation fault signal terminated program cc1
```

**Solution:** Use Clang instead of GCC. This is a known bug in GCC 15.2.1 when compiling GTK3 CGO bindings.

**Always compile with:**
```bash
CC=clang CGO_ENABLED=1 go build -o moon-phase
```

## Project Structure

```
moon-phase/
├── main.go           # Main application code
├── go.mod            # Go module dependencies
├── go.sum            # Dependency checksums
├── build.sh          # Build script (uses Clang)
├── README.md         # This file
├── SESSION.md        # Project development log
└── moon-phase        # Compiled binary
```

## Development

### Code Structure

- `julianDay()` - Converts dates to Julian Day Number
- `calculateMoonPhase()` - Calculates current moon phase
- `getPhaseName()` - Maps lunar age to phase names
- `getWiccaMeaning()` - Returns Wiccan correspondence for phase
- `drawMoon()` - Renders moon graphic with Cairo
- `main()` - GTK3 GUI setup and event handling

### Adding New Features

The code is modular and easy to extend:
- Add new spiritual traditions in separate functions like `getWiccaMeaning()`
- Extend GUI by adding labels to the vbox container
- Modify moon rendering in `drawMoon()` function

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for:
- Additional magical tradition correspondences
- UI improvements
- Localization/translations
- Bug fixes
- Documentation improvements

## Spiritual Traditions

This application uses modern Wiccan correspondences as the default spiritual framework. The system is designed to be extensible for other traditions (Traditional Witchcraft, Ceremonial Magic, etc.).

## License

This project is licensed under the MIT License - see below for details:

```
MIT License

Copyright (c) 2025 Moon Phase Contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## Acknowledgments

- [gotk3](https://github.com/gotk3/gotk3) - Go bindings for GTK3
- Astronomical calculations based on Jean Meeus' algorithms
- Wiccan correspondences from modern Wiccan tradition

## Support

For issues, questions, or suggestions:
- Open an issue on GitHub
- Check existing issues for solutions
- Review `SESSION.md` for development notes

## Changelog

### Version 1.1 (2025-12-17)
- ✨ Added Wiccan magical correspondences for each moon phase
- 📏 Increased window height to 550px to accommodate new information
- 🎨 Added italic styling for spiritual meanings

### Version 1.0 (2025-12-17)
- 🎉 Initial release
- 🌙 Accurate moon phase calculation with Julian Day algorithm
- 🎨 Cairo graphics rendering
- 🖥️ GTK3 native interface
- 📊 Detailed astronomical information display
- 🔧 Clang compilation workaround for GCC bug

---

**Made with 🌙 and ✨**
