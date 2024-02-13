const jwtToken = "eyJhbGciOiJIUzI1NiJ9.e30.XmNK3GpH3Ys_7wsYBfq4C3M6goz71I7dTgUkuIa5lyQ"

export async function findUser(token){
    const response = await fetch(`http://localhost:3000/service/find/user/${token}`);
    if(response.ok){
        const data = await response.json();
        const user = await data.data;
        return user;
    } else {
        return false;
    }
}

export async function createRoom(form){
    console.log(JSON.stringify(form))
    const room = await fetch("http://localhost:3000/service/create/room", {
        method:"POST",
        mode:"cors",
        credentials: 'include',
        body: JSON.stringify(form),
        headers: {
            'Content-Type': 'application/json',
        }
    });
    if(room.ok){
        const response = await room.json()
        const data = await response.roomID;
        window.location.href = `/room/${data}`
    }
}

export async function join(members, member, form) {
    console.log(members);
    console.log(member);

    if (members && Array.isArray(members)) { // Null kontrolü
        console.log(members.find(user => user.UserID === member.id));
        
        if (!members.find(user => user.UserID === member.id)) {
            const response = await fetch("http://localhost:3000/service/join/room", {
                method: "POST",
                mode: "cors",
                credentials: 'include',
                body: JSON.stringify(form),
                headers: {
                    'Content-Type': 'application/json',
                }
            });

            if (response.ok) {
                console.log("Kullanıcı odaya giriş yaptı");
            } else {
                console.error("Odaya katılırken bir hata oluştu");
            }
        } else {
            console.log("Kullanıcı zaten odaya katılmış durumda");
        }
    } else {
        console.error("Members dizisi boş veya tanımlanmamış.");
    }
}

export async function leave(members,member,form){
    const user = await fetch("http://localhost:3000/service/leave/room",{
        method:"POST",
        mode:"cors",
        credentials: 'include',
        body: JSON.stringify(form),
        headers: {
            'Content-Type': 'application/json',
        }
    });

    if(user.ok){
        console.log("Kullanıcı odaya giriş yaptı")
    }
}

export async function youtubeSearch(title){
    const response = await fetch(`http://localhost:5000/search/youtube?s=${title}`, {
        mode:"cors"
    });
    if(response.ok){
        const data = await response.json();
        return data.result;
    }
}