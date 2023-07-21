import type { Actions } from "@sveltejs/kit";

export const actions = {

    createChat: async ({ request }) => {
        // will create a new chat in database and return uuid
        const data = await request.formData();
        const name = data.get("name");
        const description = data.get("description");
        const userId = data.get("userID");

        // create a request to api to get new chat uuid
        let chat =
            await fetch(`http://localhost:8081/chat`, {
                method: "POST",
                body: JSON.stringify({ "name": name, "description": description }),
            }).then(res => res.json());

        let link = await fetch(`http://localhost:8081/chats/link`, {
            method: "POST",
            body: JSON.stringify({ "chat_id": chat.id, "user_id": userId }),
        }).then(res => res.status);

        return { chat_id: chat.id, name: chat.name, description: chat.description, success: link == 200 };
    },

    joinChat: async ({ request }) => {
        const data = await request.formData();
        const chatId = data.get("uuid");
        const userId = data.get("userID");

        // verify the chat exists
        const chat = await fetch(`http://localhost:8081/chats/${chatId}`).then(res => res.json());
        if (chat.created_at == undefined) {
            return { success: false };
        } else {
            const link = await fetch(`http://localhost:8081/chats/link`, {
                method: "POST",
                body: JSON.stringify({ "chat_id": chatId, "user_id": userId }),
            })
                .then(res => res.status);
            return { chat_id: chatId, name: chat.name, description: chat.description, success: link == 200 };
        }
    }

} satisfies Actions;
