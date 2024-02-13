<script>
    import {page} from '$app/stores'
    import { onMount, onDestroy } from 'svelte';
    import { initializeWebSocket, sendMessage, sendVideoRequest, changeVideo, messages, videos, users, getEmbedUrl} from '$lib/ws.js';
    import { join, leave, youtubeSearch } from '$lib/request.js'
    import Cookies from "js-cookie";
    let searchType = "embedURLs";
    let search;
    let results = [];
    let token = Cookies.get("token");
    let roomID = $page.params.id;
    let roomName;

    let username = token;
    let room = roomID;
    let message = '';

    let roomSettings;
    let user;
    let members;

    onMount(async () => {
        const roomData = await fetch(`http://localhost:3000/service/find/room/${roomID}`)
        if(roomData.ok){
            const response = await roomData.json();

            roomName = await response.data.Name;
            roomSettings = await response.data;

            setTimeout(async() => {
                const userData = await fetch(`http://localhost:3000/service/find/user/${token}`)
                if(userData.ok){
                    const response = await userData.json();

                    user = await response.data;

                    document.getElementById("video").src = "https://www.youtube.com/embed/eHzoTLwx01E";
                    initializeWebSocket(room, username);
                    setTimeout(async()=>{
                        const membersData = await fetch(`http://localhost:3000/service/members/${roomID}`)
                        if(membersData.ok){
                            const response = await membersData.json()
                            members = response.data;

                            const JoinForm = {
                                roomID   : roomID,
                                userID   : user.id,
                                username : user.username,
                                perm     : "user" 
                            }
                   
                            join(members,user,JoinForm)

                            users.set(members)
                        }
                    },500)
                } else {
                    window.location.href = "/404"
                }
            }, 10);
        } else {
            window.location.href = "/404"
        }
        window.addEventListener('beforeunload', ()=>{
            const LeaveForm = {
                roomID   : roomID,
                userID   : user.id,
            }
            leave(LeaveForm)
        });
    });

    async function searchVideo(){
        results = await youtubeSearch(document.getElementById("videoUrl").value)
    }

    function findMember(id){
        return members.find(member => member.UserID === id)
    }

    function sendMessageHandler() {
        sendMessage(username, message);
        message = '';
    }

    function sendVideoRequestHandler(url) {
        if(roomSettings.VideoRequest || findMember(user.ID).Perm == "mod" || findMember(user.ID).Perm == "admin"){
            if(searchType == "embedURL"){
                if(roomSettings.VideoRequest || findMember(user.ID).Perm == "mod" || findMember(user.ID).Perm == "admin"){
                    sendVideoRequest(username, document.getElementById("videoUrl").value);
                }
            } else {
                sendVideoRequest(username, url);
                results = []
            }
        }
    }

    function changeVideoHandler(url, id) {
        if(roomSettings.VideoRequest || findMember(user.ID).Perm == "mod" || findMember(user.ID).Perm == "admin"){
            changeVideo(username, url, id)
        }
    }

    function changeSearchType(type){
        searchType = type;
    }
</script>
    <div class="bg-surface-800 pt-2 flex items-center w-full">
        <div class="btn-group bg-surface-800">
        <button on:click={()=>{changeSearchType("youtube")}}>Youtube</button> 
        <button on:click={()=>{changeSearchType("embedURL")}}>Embed URL</button>
        </div>
    </div>
    <div class="bg-surface-800 p-2 flex items-center justify-center w-full">
        {#if searchType == "embedURL"}
            <input class="input w-[85%] p-2 rounded-r-none" type="text" id="videoUrl" placeholder="Video URL">
            <button class="btn variant-filled-primary rounded-l-none" on:click={sendVideoRequestHandler}>Add</button>
        {:else}
            <input class="input w-[85%] p-2 rounded-r-none" type="text" id="videoUrl" placeholder="Video URL">
            <button class="btn variant-filled-primary rounded-l-none" on:click={searchVideo}>Search</button>
        {/if}
    </div>
    <div class="bg-surface-800 p-2 flex items-center justify-center gap-5 flex-nowrap overflow-x-scroll max-w-[100vw]">
        {#each results as video}
            <div class="card w-[calc(25vh-5px)] min-w-[25vh] p-2">
                <iframe title="requested video" class="rounded-2xl" width="100%" style="border:none" src={getEmbedUrl(video)}></iframe>
                <br><button class="btn variant-filled-primary" on:click={()=>{sendVideoRequestHandler(getEmbedUrl(video))}} style="width:100%">Listeye ekle</button>
            </div>
        {/each}
    </div>
    <div class="items-center justify-center w-full">
        <h1 class="pl-[calc(6dvw+0.5rem)] pt-2 text-2xl">Room {roomName}</h1>
        <div class="flex pl-[6dvw] pr-[6dvw]">
            <div class="w-[60%] p-2">
                <iframe title="video" class="rounded-2xl" style="border:none;width:100%;height:50dvh" id="video"></iframe>
            </div>
            <div class="w-[40%] p-2">
                <div id="chat">
                    {#each $messages as { username, message }}
                        <div style="display:flex;align-items:center">
                            <strong style="margin-left:10px">{username}:</strong> {message}
                        </div>
                    {/each}
                </div>
                <div class="flex gap-2">
                    <input class="input p-2" type="text" bind:value={message} placeholder="Type your message">
                    <button class="btn variant-filled-primary" on:click={sendMessageHandler}>Send</button>
                </div>
            </div>
        </div>
        <div class="bg-surface-800 p-2 flex items-center justify-center gap-5 flex-nowrap overflow-x-scroll max-w-[calc(100vw-24px)]">
            {#each $videos as { id, url, by }}
                <div class="card w-[calc(25vh-5px)] min-w-[25vh] p-2">
                    <iframe title="requested video" class="rounded-2xl" width="100%" style="border:none" src={url}></iframe>
                    <br><b>{by}</b> tarafından listeye eklendi
                    <br><button class="btn variant-filled-primary" on:click={() => changeVideoHandler(url, id)} style="width:100%">Videoyu aç</button>
                </div>
            {/each}
        </div>
    </div>

    {#each $users as {Username}}
        {Username}
    {/each}
<style>
    #chat{height:45dvh;max-height:45dvh;overflow-y:scroll;position:relative;}
</style>