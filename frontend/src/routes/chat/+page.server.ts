import type { Actions } from "@sveltejs/kit";

export const actions = {

    createChat: async ({ request }) => {
        // will create a new chat in database and return uuid
        const data = await request.formData();
        const name = data.get("name");
        const description = data.get("description");

        // create a request to api to get new chat uuid
        let chat = await fetch(`http://localhost:8081/createChat?name=${name}&description=${description}`).then(res => res.json());

        let link = await fetch(`http://localhost:8081/linkChatAndUser?chat=${chat.id}&user=${data.get("userID")}`)

        return { uuid: chat.id, name: chat.name, description: chat.description, success: link.status };
    },

    joinChat: async ({ request }) => {
        const data = await request.formData();
        const chatId = data.get("uuid");
        const userId = data.get("userID");

        let res = await fetch(`http://localhost:8081/linkChatAndUser?chat=${chatId}&user=${userId}`)
        if (res.status == 200) {

            const chatData = await fetch(`http://localhost:8081/checkChat?id=${chatId}`).then(res => res.json());
            console.log(chatData)

            return { success: true, uuid: chatData.id, name: chatData.name, description: chatData.description }
        } else {
            return { success: false }
        }
    }

} satisfies Actions;
