# TODO

## Ludum Dare Priorities

- [x] CTRL+S to Save
    - [x] Add a callback for saving
    - [x] Modify editors to allow saving
    - [x] Introduce file handles and saving to SpriteSets
- [x] Clone the Default SpriteSet when creating a new set or starting from nothing
    - [x] Clone for SpriteSet
    - [x] Clone for Sprite
- [x] Add Button type for sending signals to other UI elements
- [ ] Modify Sprite Sets
    - [ ] Add Sprite
    - [ ] Duplicate Sprite
    - [ ] Copy Sprite
    - [ ] Clear Sprite
    - [ ] Remove Sprite
    - [ ] Select Sprite
        - [ ] Switching Sprites using Selector
    - [ ] Reorder Sprites
        - [ ] First
        - [ ] Left
        - [ ] Right
        - [ ] Last

- [x] Toolbar
    - [x] Working Save function
- [x] Load Sprite-Set from os.Args[1]

## Post LD priorities

- [ ] Embed assets using embed.FS for the UI
    - [ ] Toolbar sprites
    - [ ] Fonts
    - [ ] Default Palettes
- [ ] Are you sure?
    - [ ] Modal Dialog
    - [ ] Translucent Overlay
- [ ] Encode Sprites as a Bitmap-like format
    - [ ] Custom image.Image types
        - [ ] 1-bit per pixel (2 colors)
        - [ ] 2-bit per pixel (4 colors)
        - [ ] 4-bit per pixel (16 colors)
        - [ ] 8-bit per pixel (256 colors) (existing sprite encoding)
    - [ ] Update Sprite Set to support optimal sizes

## General

- [ ] Color picker
- [ ] Animated loading screen with gameboy inspired start-up sound

## Serialization

**Sprites**

- [x] Load Sprite Set from JSON
- [ ] Save Sprite Set to JSON

**Palettes**

- [x] Load Palette from JSON
- [ ] Save Palette to JSON

## Data Sharing

**Sprites**

 - [ ] Loaded Sprites are available in Sprite Editor, Sprite Animator, Palette Editor, and Map Editor
   - [ ] Last Selected Sprite in Sprite Editor used as Preview in Palette Editor
   - [ ] Loading a new Sprite Set will clear the Animation Editor (if user confirms load)
   - [ ] Loading a new Sprite Set will update the Map Editor (if user confirms load)

**Animations**

 - [ ] Loaded Animation is only available in the Animation Editor
 - [ ] Loading a new Sprite Set will clear the Animation Editor (if user confirms load)

**Palettes**

 - [ ] Loaded Palettes are available in Sprite Editor, Palette Editor, and Map Editor

**Maps**

 - [ ] Loaded Map is only available in the Map Editor
 - [ ] Loading a new Sprite Set will update the Map Editor (if user confirms load)

## Sprite Editor

### Toolbar

**File Operations**

 - [ ] New
 - [ ] Open
 - [ ] Save
 - [ ] Save As...
 - [ ] Export Bitmap

**Sprite Operations**

 - [ ] Add
 - [ ] Copy
 - [ ] Remove

### Sprite Selection

 - **Note:** Consider limiting the number of sprites in a set and using pages to switch sets
 - [ ] Grid
   - [ ] Switch Sprite
   - [ ] Add Sprite (+)
 - [ ] Prev/Next Page
 - [ ] 2x2 Tile Preview

### Editor

**Sprite Canvas**

 - [x] Left-Click Set Color
 - [ ] Middle-Click Replace Color
 - [x] Right-Click Clear Color (Transparent)
 - [x] Live Preview

**Color Chooser**

 - [ ] Scroll-Wheel Switch Color
 - [ ] Left-Click Select Color
 - [ ] Prev/Next Palette

## Sprite Animator

### Toolbar

**File Operations**

 - [ ] New
 - [ ] Open
 - [ ] Save
 - [ ] Save As...
 - [ ] Export Bitmap
 - [ ] Export GIF

### Sprite Chooser

 - [ ] Borrow From Sprite Editor

### Keyframe Selection

 - [ ] Add Sprite
 - [ ] Insert Sprite
 - [ ] Remove Sprite
 - [ ] Replace Sprite
 - [ ] Move Up/Down

### Animation Preview

 - [ ] Preview Window
 - [ ] Speed Increase / Decrease

## Palette Editor

### Toolbar

**File Operations**

 - [ ] New
 - [ ] Open
 - [ ] Save
 - [ ] Save As...
 - [ ] Export XML

### Palette Selection

 - [ ] Add Color
 - [ ] Insert Color
 - [ ] Remove Color
 - [ ] Select Color
 - [ ] Move Up/Down

 - [ ] Add Palette
 - [ ] Remove Palette
 - [ ] Prev/Next Palette

### Color Chooser

 - [ ] RGB
 - [ ] HSV

### Preview

 - [ ] Apply Palette to last selected Sprite from Sprite Editor

## Map Editor

### Toolbar

**File Operations**

 - [ ] New
 - [ ] Open
 - [ ] Save
 - [ ] Save As...
 - [ ] Export XML

### Tile Selector

 - [ ] Borrow from Sprite Editor

### Palette Chooser

 - [ ] Adapt from Sprite Editor

### Map

 - [ ] Add/Remove Column
 - [ ] Add/Remove Row
 - [ ] Zoom In/Out
   - [ ] Set/Clear Tile
