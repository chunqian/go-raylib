// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Thu, 22 Apr 2021 14:00:10 CST.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package raylib

/*
#include "../lib/raygui/src/raygui.h"
#include "../lib/raygui/src/ricons.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// RAYGUI_VERSION as defined in src/raygui.h:134
	RAYGUI_VERSION = "2.6-dev"
	// RAYGUIDEF as defined in src/raygui.h:154
	RAYGUIDEF = 0
	// NUM_CONTROLS as defined in src/raygui.h:176
	NUM_CONTROLS = 16
	// NUM_PROPS_DEFAULT as defined in src/raygui.h:177
	NUM_PROPS_DEFAULT = 16
	// NUM_PROPS_EXTENDED as defined in src/raygui.h:178
	NUM_PROPS_EXTENDED = 8
	// TEXTEDIT_CURSOR_BLINK_FRAMES as defined in src/raygui.h:180
	TEXTEDIT_CURSOR_BLINK_FRAMES = 20
	// RICON_MAX_ICONS as defined in src/ricons.h:18
	RICON_MAX_ICONS = 256
	// RICON_SIZE as defined in src/ricons.h:19
	RICON_SIZE = 16
	// RICON_MAX_NAME_LENGTH as defined in src/ricons.h:21
	RICON_MAX_NAME_LENGTH = 32
	// RICON_DATA_ELEMENTS as defined in src/ricons.h:27
	RICON_DATA_ELEMENTS = (RICON_SIZE * RICON_SIZE / 32)
)

// GuiControlState as declared in src/raygui.h:264
type GuiControlState int32

// GuiControlState enumeration from src/raygui.h:264
const (
	GUI_STATE_NORMAL   GuiControlState = iota
	GUI_STATE_FOCUSED  GuiControlState = 1
	GUI_STATE_PRESSED  GuiControlState = 2
	GUI_STATE_DISABLED GuiControlState = 3
)

// GuiTextAlignment as declared in src/raygui.h:271
type GuiTextAlignment int32

// GuiTextAlignment enumeration from src/raygui.h:271
const (
	GUI_TEXT_ALIGN_LEFT   GuiTextAlignment = iota
	GUI_TEXT_ALIGN_CENTER GuiTextAlignment = 1
	GUI_TEXT_ALIGN_RIGHT  GuiTextAlignment = 2
)

// GuiControl as declared in src/raygui.h:291
type GuiControl int32

// GuiControl enumeration from src/raygui.h:291
const (
	DEFAULT     GuiControl = iota
	LABEL       GuiControl = 1
	BUTTON      GuiControl = 2
	TOGGLE      GuiControl = 3
	SLIDER      GuiControl = 4
	PROGRESSBAR GuiControl = 5
	CHECKBOX    GuiControl = 6
	COMBOBOX    GuiControl = 7
	DROPDOWNBOX GuiControl = 8
	TEXTBOX     GuiControl = 9
	VALUEBOX    GuiControl = 10
	SPINNER     GuiControl = 11
	LISTVIEW    GuiControl = 12
	COLORPICKER GuiControl = 13
	SCROLLBAR   GuiControl = 14
	STATUSBAR   GuiControl = 15
)

// GuiControlProperty as declared in src/raygui.h:311
type GuiControlProperty int32

// GuiControlProperty enumeration from src/raygui.h:311
const (
	BORDER_COLOR_NORMAL   GuiControlProperty = iota
	BASE_COLOR_NORMAL     GuiControlProperty = 1
	TEXT_COLOR_NORMAL     GuiControlProperty = 2
	BORDER_COLOR_FOCUSED  GuiControlProperty = 3
	BASE_COLOR_FOCUSED    GuiControlProperty = 4
	TEXT_COLOR_FOCUSED    GuiControlProperty = 5
	BORDER_COLOR_PRESSED  GuiControlProperty = 6
	BASE_COLOR_PRESSED    GuiControlProperty = 7
	TEXT_COLOR_PRESSED    GuiControlProperty = 8
	BORDER_COLOR_DISABLED GuiControlProperty = 9
	BASE_COLOR_DISABLED   GuiControlProperty = 10
	TEXT_COLOR_DISABLED   GuiControlProperty = 11
	BORDER_WIDTH          GuiControlProperty = 12
	TEXT_PADDING          GuiControlProperty = 13
	TEXT_ALIGNMENT        GuiControlProperty = 14
	RESERVED              GuiControlProperty = 15
)

// GuiDefaultProperty as declared in src/raygui.h:322
type GuiDefaultProperty int32

// GuiDefaultProperty enumeration from src/raygui.h:322
const (
	TEXT_SIZE        GuiDefaultProperty = 16
	TEXT_SPACING     GuiDefaultProperty = 17
	LINE_COLOR       GuiDefaultProperty = 18
	BACKGROUND_COLOR GuiDefaultProperty = 19
)

// GuiToggleProperty as declared in src/raygui.h:333
type GuiToggleProperty int32

// GuiToggleProperty enumeration from src/raygui.h:333
const (
	GROUP_PADDING GuiToggleProperty = 16
)

// GuiSliderProperty as declared in src/raygui.h:339
type GuiSliderProperty int32

// GuiSliderProperty enumeration from src/raygui.h:339
const (
	SLIDER_WIDTH   GuiSliderProperty = 16
	SLIDER_PADDING GuiSliderProperty = 17
)

// GuiProgressBarProperty as declared in src/raygui.h:344
type GuiProgressBarProperty int32

// GuiProgressBarProperty enumeration from src/raygui.h:344
const (
	PROGRESS_PADDING GuiProgressBarProperty = 16
)

// GuiCheckBoxProperty as declared in src/raygui.h:349
type GuiCheckBoxProperty int32

// GuiCheckBoxProperty enumeration from src/raygui.h:349
const (
	CHECK_PADDING GuiCheckBoxProperty = 16
)

// GuiComboBoxProperty as declared in src/raygui.h:355
type GuiComboBoxProperty int32

// GuiComboBoxProperty enumeration from src/raygui.h:355
const (
	COMBO_BUTTON_WIDTH   GuiComboBoxProperty = 16
	COMBO_BUTTON_PADDING GuiComboBoxProperty = 17
)

// GuiDropdownBoxProperty as declared in src/raygui.h:361
type GuiDropdownBoxProperty int32

// GuiDropdownBoxProperty enumeration from src/raygui.h:361
const (
	ARROW_PADDING          GuiDropdownBoxProperty = 16
	DROPDOWN_ITEMS_PADDING GuiDropdownBoxProperty = 17
)

// GuiTextBoxProperty as declared in src/raygui.h:369
type GuiTextBoxProperty int32

// GuiTextBoxProperty enumeration from src/raygui.h:369
const (
	TEXT_INNER_PADDING GuiTextBoxProperty = 16
	TEXT_LINES_PADDING GuiTextBoxProperty = 17
	COLOR_SELECTED_FG  GuiTextBoxProperty = 18
	COLOR_SELECTED_BG  GuiTextBoxProperty = 19
)

// GuiSpinnerProperty as declared in src/raygui.h:375
type GuiSpinnerProperty int32

// GuiSpinnerProperty enumeration from src/raygui.h:375
const (
	SPIN_BUTTON_WIDTH   GuiSpinnerProperty = 16
	SPIN_BUTTON_PADDING GuiSpinnerProperty = 17
)

// GuiScrollBarProperty as declared in src/raygui.h:385
type GuiScrollBarProperty int32

// GuiScrollBarProperty enumeration from src/raygui.h:385
const (
	ARROWS_SIZE           GuiScrollBarProperty = 16
	ARROWS_VISIBLE        GuiScrollBarProperty = 17
	SCROLL_SLIDER_PADDING GuiScrollBarProperty = 18
	SCROLL_SLIDER_SIZE    GuiScrollBarProperty = 19
	SCROLL_PADDING        GuiScrollBarProperty = 20
	SCROLL_SPEED          GuiScrollBarProperty = 21
)

// GuiScrollBarSide as declared in src/raygui.h:391
type GuiScrollBarSide int32

// GuiScrollBarSide enumeration from src/raygui.h:391
const (
	SCROLLBAR_LEFT_SIDE  GuiScrollBarSide = iota
	SCROLLBAR_RIGHT_SIDE GuiScrollBarSide = 1
)

// GuiListViewProperty as declared in src/raygui.h:399
type GuiListViewProperty int32

// GuiListViewProperty enumeration from src/raygui.h:399
const (
	LIST_ITEMS_HEIGHT  GuiListViewProperty = 16
	LIST_ITEMS_PADDING GuiListViewProperty = 17
	SCROLLBAR_WIDTH    GuiListViewProperty = 18
	SCROLLBAR_SIDE     GuiListViewProperty = 19
)

// GuiColorPickerProperty as declared in src/raygui.h:408
type GuiColorPickerProperty int32

// GuiColorPickerProperty enumeration from src/raygui.h:408
const (
	COLOR_SELECTOR_SIZE      GuiColorPickerProperty = 16
	HUEBAR_WIDTH             GuiColorPickerProperty = 17
	HUEBAR_PADDING           GuiColorPickerProperty = 18
	HUEBAR_SELECTOR_HEIGHT   GuiColorPickerProperty = 19
	HUEBAR_SELECTOR_OVERFLOW GuiColorPickerProperty = 20
)

// GuiIconName as declared in src/ricons.h:289
type GuiIconName int32

// GuiIconName enumeration from src/ricons.h:289
const (
	RICON_NONE                    GuiIconName = iota
	RICON_FOLDER_FILE_OPEN        GuiIconName = 1
	RICON_FILE_SAVE_CLASSIC       GuiIconName = 2
	RICON_FOLDER_OPEN             GuiIconName = 3
	RICON_FOLDER_SAVE             GuiIconName = 4
	RICON_FILE_OPEN               GuiIconName = 5
	RICON_FILE_SAVE               GuiIconName = 6
	RICON_FILE_EXPORT             GuiIconName = 7
	RICON_FILE_NEW                GuiIconName = 8
	RICON_FILE_DELETE             GuiIconName = 9
	RICON_FILETYPE_TEXT           GuiIconName = 10
	RICON_FILETYPE_AUDIO          GuiIconName = 11
	RICON_FILETYPE_IMAGE          GuiIconName = 12
	RICON_FILETYPE_PLAY           GuiIconName = 13
	RICON_FILETYPE_VIDEO          GuiIconName = 14
	RICON_FILETYPE_INFO           GuiIconName = 15
	RICON_FILE_COPY               GuiIconName = 16
	RICON_FILE_CUT                GuiIconName = 17
	RICON_FILE_PASTE              GuiIconName = 18
	RICON_CURSOR_HAND             GuiIconName = 19
	RICON_CURSOR_POINTER          GuiIconName = 20
	RICON_CURSOR_CLASSIC          GuiIconName = 21
	RICON_PENCIL                  GuiIconName = 22
	RICON_PENCIL_BIG              GuiIconName = 23
	RICON_BRUSH_CLASSIC           GuiIconName = 24
	RICON_BRUSH_PAINTER           GuiIconName = 25
	RICON_WATER_DROP              GuiIconName = 26
	RICON_COLOR_PICKER            GuiIconName = 27
	RICON_RUBBER                  GuiIconName = 28
	RICON_COLOR_BUCKET            GuiIconName = 29
	RICON_TEXT_T                  GuiIconName = 30
	RICON_TEXT_A                  GuiIconName = 31
	RICON_SCALE                   GuiIconName = 32
	RICON_RESIZE                  GuiIconName = 33
	RICON_FILTER_POINT            GuiIconName = 34
	RICON_FILTER_BILINEAR         GuiIconName = 35
	RICON_CROP                    GuiIconName = 36
	RICON_CROP_ALPHA              GuiIconName = 37
	RICON_SQUARE_TOGGLE           GuiIconName = 38
	RICON_SYMMETRY                GuiIconName = 39
	RICON_SYMMETRY_HORIZONTAL     GuiIconName = 40
	RICON_SYMMETRY_VERTICAL       GuiIconName = 41
	RICON_LENS                    GuiIconName = 42
	RICON_LENS_BIG                GuiIconName = 43
	RICON_EYE_ON                  GuiIconName = 44
	RICON_EYE_OFF                 GuiIconName = 45
	RICON_FILTER_TOP              GuiIconName = 46
	RICON_FILTER                  GuiIconName = 47
	RICON_TARGET_POINT            GuiIconName = 48
	RICON_TARGET_SMALL            GuiIconName = 49
	RICON_TARGET_BIG              GuiIconName = 50
	RICON_TARGET_MOVE             GuiIconName = 51
	RICON_CURSOR_MOVE             GuiIconName = 52
	RICON_CURSOR_SCALE            GuiIconName = 53
	RICON_CURSOR_SCALE_RIGHT      GuiIconName = 54
	RICON_CURSOR_SCALE_LEFT       GuiIconName = 55
	RICON_UNDO                    GuiIconName = 56
	RICON_REDO                    GuiIconName = 57
	RICON_REREDO                  GuiIconName = 58
	RICON_MUTATE                  GuiIconName = 59
	RICON_ROTATE                  GuiIconName = 60
	RICON_REPEAT                  GuiIconName = 61
	RICON_SHUFFLE                 GuiIconName = 62
	RICON_EMPTYBOX                GuiIconName = 63
	RICON_TARGET                  GuiIconName = 64
	RICON_TARGET_SMALL_FILL       GuiIconName = 65
	RICON_TARGET_BIG_FILL         GuiIconName = 66
	RICON_TARGET_MOVE_FILL        GuiIconName = 67
	RICON_CURSOR_MOVE_FILL        GuiIconName = 68
	RICON_CURSOR_SCALE_FILL       GuiIconName = 69
	RICON_CURSOR_SCALE_RIGHT_FILL GuiIconName = 70
	RICON_CURSOR_SCALE_LEFT_FILL  GuiIconName = 71
	RICON_UNDO_FILL               GuiIconName = 72
	RICON_REDO_FILL               GuiIconName = 73
	RICON_REREDO_FILL             GuiIconName = 74
	RICON_MUTATE_FILL             GuiIconName = 75
	RICON_ROTATE_FILL             GuiIconName = 76
	RICON_REPEAT_FILL             GuiIconName = 77
	RICON_SHUFFLE_FILL            GuiIconName = 78
	RICON_EMPTYBOX_SMALL          GuiIconName = 79
	RICON_BOX                     GuiIconName = 80
	RICON_BOX_TOP                 GuiIconName = 81
	RICON_BOX_TOP_RIGHT           GuiIconName = 82
	RICON_BOX_RIGHT               GuiIconName = 83
	RICON_BOX_BOTTOM_RIGHT        GuiIconName = 84
	RICON_BOX_BOTTOM              GuiIconName = 85
	RICON_BOX_BOTTOM_LEFT         GuiIconName = 86
	RICON_BOX_LEFT                GuiIconName = 87
	RICON_BOX_TOP_LEFT            GuiIconName = 88
	RICON_BOX_CENTER              GuiIconName = 89
	RICON_BOX_CIRCLE_MASK         GuiIconName = 90
	RICON_POT                     GuiIconName = 91
	RICON_ALPHA_MULTIPLY          GuiIconName = 92
	RICON_ALPHA_CLEAR             GuiIconName = 93
	RICON_DITHERING               GuiIconName = 94
	RICON_MIPMAPS                 GuiIconName = 95
	RICON_BOX_GRID                GuiIconName = 96
	RICON_GRID                    GuiIconName = 97
	RICON_BOX_CORNERS_SMALL       GuiIconName = 98
	RICON_BOX_CORNERS_BIG         GuiIconName = 99
	RICON_FOUR_BOXES              GuiIconName = 100
	RICON_GRID_FILL               GuiIconName = 101
	RICON_BOX_MULTISIZE           GuiIconName = 102
	RICON_ZOOM_SMALL              GuiIconName = 103
	RICON_ZOOM_MEDIUM             GuiIconName = 104
	RICON_ZOOM_BIG                GuiIconName = 105
	RICON_ZOOM_ALL                GuiIconName = 106
	RICON_ZOOM_CENTER             GuiIconName = 107
	RICON_BOX_DOTS_SMALL          GuiIconName = 108
	RICON_BOX_DOTS_BIG            GuiIconName = 109
	RICON_BOX_CONCENTRIC          GuiIconName = 110
	RICON_BOX_GRID_BIG            GuiIconName = 111
	RICON_OK_TICK                 GuiIconName = 112
	RICON_CROSS                   GuiIconName = 113
	RICON_ARROW_LEFT              GuiIconName = 114
	RICON_ARROW_RIGHT             GuiIconName = 115
	RICON_ARROW_BOTTOM            GuiIconName = 116
	RICON_ARROW_TOP               GuiIconName = 117
	RICON_ARROW_LEFT_FILL         GuiIconName = 118
	RICON_ARROW_RIGHT_FILL        GuiIconName = 119
	RICON_ARROW_BOTTOM_FILL       GuiIconName = 120
	RICON_ARROW_TOP_FILL          GuiIconName = 121
	RICON_AUDIO                   GuiIconName = 122
	RICON_FX                      GuiIconName = 123
	RICON_WAVE                    GuiIconName = 124
	RICON_WAVE_SINUS              GuiIconName = 125
	RICON_WAVE_SQUARE             GuiIconName = 126
	RICON_WAVE_TRIANGULAR         GuiIconName = 127
	RICON_CROSS_SMALL             GuiIconName = 128
	RICON_PLAYER_PREVIOUS         GuiIconName = 129
	RICON_PLAYER_PLAY_BACK        GuiIconName = 130
	RICON_PLAYER_PLAY             GuiIconName = 131
	RICON_PLAYER_PAUSE            GuiIconName = 132
	RICON_PLAYER_STOP             GuiIconName = 133
	RICON_PLAYER_NEXT             GuiIconName = 134
	RICON_PLAYER_RECORD           GuiIconName = 135
	RICON_MAGNET                  GuiIconName = 136
	RICON_LOCK_CLOSE              GuiIconName = 137
	RICON_LOCK_OPEN               GuiIconName = 138
	RICON_CLOCK                   GuiIconName = 139
	RICON_TOOLS                   GuiIconName = 140
	RICON_GEAR                    GuiIconName = 141
	RICON_GEAR_BIG                GuiIconName = 142
	RICON_BIN                     GuiIconName = 143
	RICON_HAND_POINTER            GuiIconName = 144
	RICON_LASER                   GuiIconName = 145
	RICON_COIN                    GuiIconName = 146
	RICON_EXPLOSION               GuiIconName = 147
	RICON_1UP                     GuiIconName = 148
	RICON_PLAYER                  GuiIconName = 149
	RICON_PLAYER_JUMP             GuiIconName = 150
	RICON_KEY                     GuiIconName = 151
	RICON_DEMON                   GuiIconName = 152
	RICON_TEXT_POPUP              GuiIconName = 153
	RICON_GEAR_EX                 GuiIconName = 154
	RICON_CRACK                   GuiIconName = 155
	RICON_CRACK_POINTS            GuiIconName = 156
	RICON_STAR                    GuiIconName = 157
	RICON_DOOR                    GuiIconName = 158
	RICON_EXIT                    GuiIconName = 159
	RICON_MODE_2D                 GuiIconName = 160
	RICON_MODE_3D                 GuiIconName = 161
	RICON_CUBE                    GuiIconName = 162
	RICON_CUBE_FACE_TOP           GuiIconName = 163
	RICON_CUBE_FACE_LEFT          GuiIconName = 164
	RICON_CUBE_FACE_FRONT         GuiIconName = 165
	RICON_CUBE_FACE_BOTTOM        GuiIconName = 166
	RICON_CUBE_FACE_RIGHT         GuiIconName = 167
	RICON_CUBE_FACE_BACK          GuiIconName = 168
	RICON_CAMERA                  GuiIconName = 169
	RICON_SPECIAL                 GuiIconName = 170
	RICON_LINK_NET                GuiIconName = 171
	RICON_LINK_BOXES              GuiIconName = 172
	RICON_LINK_MULTI              GuiIconName = 173
	RICON_LINK                    GuiIconName = 174
	RICON_LINK_BROKE              GuiIconName = 175
	RICON_TEXT_NOTES              GuiIconName = 176
	RICON_NOTEBOOK                GuiIconName = 177
	RICON_SUITCASE                GuiIconName = 178
	RICON_SUITCASE_ZIP            GuiIconName = 179
	RICON_MAILBOX                 GuiIconName = 180
	RICON_MONITOR                 GuiIconName = 181
	RICON_PRINTER                 GuiIconName = 182
	RICON_PHOTO_CAMERA            GuiIconName = 183
	RICON_PHOTO_CAMERA_FLASH      GuiIconName = 184
	RICON_HOUSE                   GuiIconName = 185
	RICON_HEART                   GuiIconName = 186
	RICON_CORNER                  GuiIconName = 187
	RICON_VERTICAL_BARS           GuiIconName = 188
	RICON_VERTICAL_BARS_FILL      GuiIconName = 189
	RICON_LIFE_BARS               GuiIconName = 190
	RICON_INFO                    GuiIconName = 191
	RICON_CROSSLINE               GuiIconName = 192
	RICON_HELP                    GuiIconName = 193
	RICON_FILETYPE_ALPHA          GuiIconName = 194
	RICON_FILETYPE_HOME           GuiIconName = 195
	RICON_LAYERS_VISIBLE          GuiIconName = 196
	RICON_LAYERS                  GuiIconName = 197
	RICON_WINDOW                  GuiIconName = 198
	RICON_HIDPI                   GuiIconName = 199
	RICON_200                     GuiIconName = 200
	RICON_201                     GuiIconName = 201
	RICON_202                     GuiIconName = 202
	RICON_203                     GuiIconName = 203
	RICON_204                     GuiIconName = 204
	RICON_205                     GuiIconName = 205
	RICON_206                     GuiIconName = 206
	RICON_207                     GuiIconName = 207
	RICON_208                     GuiIconName = 208
	RICON_209                     GuiIconName = 209
	RICON_210                     GuiIconName = 210
	RICON_211                     GuiIconName = 211
	RICON_212                     GuiIconName = 212
	RICON_213                     GuiIconName = 213
	RICON_214                     GuiIconName = 214
	RICON_215                     GuiIconName = 215
	RICON_216                     GuiIconName = 216
	RICON_217                     GuiIconName = 217
	RICON_218                     GuiIconName = 218
	RICON_219                     GuiIconName = 219
	RICON_220                     GuiIconName = 220
	RICON_221                     GuiIconName = 221
	RICON_222                     GuiIconName = 222
	RICON_223                     GuiIconName = 223
	RICON_224                     GuiIconName = 224
	RICON_225                     GuiIconName = 225
	RICON_226                     GuiIconName = 226
	RICON_227                     GuiIconName = 227
	RICON_228                     GuiIconName = 228
	RICON_229                     GuiIconName = 229
	RICON_230                     GuiIconName = 230
	RICON_231                     GuiIconName = 231
	RICON_232                     GuiIconName = 232
	RICON_233                     GuiIconName = 233
	RICON_234                     GuiIconName = 234
	RICON_235                     GuiIconName = 235
	RICON_236                     GuiIconName = 236
	RICON_237                     GuiIconName = 237
	RICON_238                     GuiIconName = 238
	RICON_239                     GuiIconName = 239
	RICON_240                     GuiIconName = 240
	RICON_241                     GuiIconName = 241
	RICON_242                     GuiIconName = 242
	RICON_243                     GuiIconName = 243
	RICON_244                     GuiIconName = 244
	RICON_245                     GuiIconName = 245
	RICON_246                     GuiIconName = 246
	RICON_247                     GuiIconName = 247
	RICON_248                     GuiIconName = 248
	RICON_249                     GuiIconName = 249
	RICON_250                     GuiIconName = 250
	RICON_251                     GuiIconName = 251
	RICON_252                     GuiIconName = 252
	RICON_253                     GuiIconName = 253
	RICON_254                     GuiIconName = 254
	RICON_255                     GuiIconName = 255
)
