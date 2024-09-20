
# Weather App

This is a simple weather application built with React and Go. The frontend is built using React, and the backend is built using Go. The app fetches weather data from the OpenWeather API and displays it to the user, including temperature, weather condition, and an icon representing the current weather.

## Features

- Fetches real-time weather data for a specified city
- Displays temperature, weather condition, and an icon representing the weather
- Error handling for invalid city names or failed API requests

## Prerequisites

- **Node.js** and **npm**
- **Go** (Go 1.16+)
- An **OpenWeatherMap API key** (You can get it from [here](https://openweathermap.org/api))

## Getting Started

### 1. Clone the Repository
\`\`\`bash
git clone https://github.com/YOUR_GITHUB_USERNAME/REPO_NAME.git
cd REPO_NAME
\`\`\`

### 2. Frontend Setup (React)

1. **Navigate to the frontend directory:**
   \`\`\`bash
   cd frontend
   \`\`\`

2. **Install dependencies:**
   \`\`\`bash
   npm install
   \`\`\`

3. **Start the frontend:**
   \`\`\`bash
   npm start
   \`\`\`

   This will start the frontend on \`http://localhost:3000\`.

### 3. Backend Setup (Go)

1. **Navigate to the backend directory:**
   \`\`\`bash
   cd cmd/server
   \`\`\`

2. **Install Go dependencies:**
   \`\`\`bash
   go get
   \`\`\`

3. **Set up your OpenWeather API key:**

   Open the \`cmd/server/main.go\` file and replace the \`apiKey\` constant with your OpenWeatherMap API key:
   \`\`\`go
   const apiKey = "YOUR_API_KEY_HERE"
   \`\`\`

4. **Start the Go server:**
   \`\`\`bash
   go run main.go
   \`\`\`

   This will start the backend server on \`http://localhost:8080\`.

## How to Use

1. In the browser, go to \`http://localhost:3000\`.
2. Enter the name of a city in the input field and click "Get Weather."
3. The app will display the city's temperature, weather condition, and an icon representing the weather.

## Next Steps

1. **Refine Weather Data:**
   - Fetch more detailed weather information (humidity, wind speed, etc.) and display it on the frontend.

2. **UI Enhancements:**
   - Improve the user interface by adding input validation, animations, and a more visually appealing design.

3. **Deploy the App:**
   - Consider deploying the frontend using services like Vercel or Netlify and the backend using services like Heroku or AWS.
