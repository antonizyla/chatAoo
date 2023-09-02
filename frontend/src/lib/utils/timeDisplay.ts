function timeDisplay(timestamp: string): string {
    // recieves a timestamp as a string 
    const time = new Date(timestamp)
    const currentTime = new Date()
    if (time.getDay() == currentTime.getDay()) {
        return time.toLocaleTimeString().slice(0, -3)
    } else if (currentTime.getDay() - time.getDay() == 1) {
        return time.toLocaleTimeString().slice(0, -3) + " Yesterday"
    }
    return time.toLocaleString().slice(0, -3)
}

export { timeDisplay };
