<script>
    import { onMount } from "svelte";
    import {AppBar} from "@skeletonlabs/skeleton";
    import Cookies from "js-cookie";
    import {findUser, createRoom} from "$lib/request.js";
    
    let user = false;

    let name;
    let description;
    let roomStatus;
    let stopVideo;
    let changeVideo;
    let videoRequest;
    
    onMount(async()=>{
        const result = await findUser(Cookies.get("token"));
        if(result){
            user = result;
        } else {
            console.log("Kullanıcı bulunamadı")
        }
    })

    function CreateRoom(){
        const formData= {
            "name": name,
            "description": description,
            "public": roomStatus,
            "stopVideo":stopVideo,
            "changeVideo": changeVideo,
            "videoRequest": videoRequest,
        }
        console.log(JSON.stringify(formData))
        createRoom(formData)
    }

</script>
<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
	<svelte:fragment slot="lead"><div class="text-2xl">Full-stack W2G</div></svelte:fragment>
	<div></div>
    <svelte:fragment slot="trail">
    {#if !user}
        <a href="/login" class="btn bg-primary-700">Login</a>
    {:else}
        <button class="btn bg-primary-700">Create Room</button>
    {/if}
    </svelte:fragment>
</AppBar>
{#if user}
    <div class="card p-4 mt-5">
        <h1 class="text-2xl">Create Room</h1>
        <input class="input p-2 mt-2" type="text" placeholder="Room name" bind:value={name}/>
        <input class="input p-2 mt-2" type="text" placeholder="Room Description" bind:value={description}/>
        <input type="checkbox" bind:checked={roomStatus}> 
        <span>Everyone can see your room?</span><br>

        <input type="checkbox" bind:checked={stopVideo}>
        <span>Users can stop video?</span><br>

        <input type="checkbox" bind:checked={changeVideo}>
        <span>Users can change video?</span><br>
        
        <input type="checkbox" bind:checked={videoRequest}>
        <span>Users can request video?</span><br>
        <button class="btn p-2 bg-primary-700" on:click={CreateRoom}>Create</button>
    </div>
{/if}