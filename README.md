# Weather API for Waybar

![Weather Module Preview]![image](https://github.com/user-attachments/assets/98a89e00-a82f-42e9-b91b-30245ddd4496)


This script fetches weather data from [WeatherAPI](https://www.weatherapi.com/) and displays it in Waybar.

## Getting Started

### 1. Get Your API Key
Register at [WeatherAPI](https://www.weatherapi.com/) to obtain your free API key.

### 2. Configure Waybar
Add the following entry to your Waybar configuration file (typically `~/.config/waybar/config.json`):

```json
"custom/weather": {
    "exec": "~/.config/waybar/scripts/weather",
    "interval": 300,
    "tooltip": false,
    "return-type": "json",
    "format": "{}",
    "style": "~/.config/waybar/style2.css",
    "format-icons": ["🌡"]
}
```

### 3. Apply Styling (Optional)
Add the following to your Waybar CSS file (typically `~/.config/waybar/style.css`) to adjust the module's spacing:

```css
#custom-weather {
    padding-top: 3px;
    margin: 0 15px;
}
```

### 4. Restart Waybar
After making these changes, restart Waybar to apply the new configuration:
```sh
pkill waybar && waybar &
```



