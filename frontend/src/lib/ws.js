import { writable } from 'svelte/store';

let socket;

export const messages = writable([]);
export const videos = writable([]);
export const users = writable([]);

let lastVideo;

export function initializeWebSocket(room, username) {
    socket = new WebSocket(`ws://localhost:3000/ws/${room}`);

    socket.addEventListener("open", (event) => {
        console.log("WebSocket connection opened:", event);
        socket.send(JSON.stringify({ type: "join", username }));
    });

    socket.addEventListener("message", async (event) => {
        const data = JSON.parse(event.data);
        console.log(data);
        switch (data.type) {
            case "message":
                messages.update(messages => [...messages, { username: data.username, message: data.message }]);
                break;
            case "change_video":
                document.getElementById("video").src = getEmbedUrl(data.url)
                videos.update(videos => videos.filter(video => video.id !== lastVideo))
            break;
            case "request_video":
                videos.update(videos => [{ id: (videos.length+1),url: getEmbedUrl(data.url), by: data.by }, ...videos]);
            break;
            case "join":
                
            break;
            case "leave":
                videos.update(users => users.filter(user => user.username !== data.username));
            break;
        }
    });

    socket.addEventListener("close", (event) => {
        console.log("WebSocket connection closed:", event);
    });

    socket.addEventListener("error", (error) => {
        console.error("WebSocket error:", error);
    });
}

export function sendMessage(username, message) {
    if (username && message) {
        socket.send(JSON.stringify({ type: "message", username, message }));
    } else {
        console.error("Username and message are required.");
    }
}

export function sendVideoRequest(username, url) {
    if (username && url) {
        socket.send(JSON.stringify({ type: "request_video", by: username, url }));
    } else {
        console.error("Video by and URL are required.");
    }
}

export function changeVideo(username, url, id){
    if (username && url) {
        lastVideo = id;
        socket.send(JSON.stringify({ type: "change_video", by: username, url }));
    } else {
        console.error("Video by and URL are required.");
    }
}

export function getEmbedUrl(url) {
    let newUrl = url.replace("watch?v=", "embed/");
    return newUrl.split("&")[0];
}
