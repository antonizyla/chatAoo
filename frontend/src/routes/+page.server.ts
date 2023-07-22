import type { Actions } from "@sveltejs/kit";

export const actions = {

    createUser: async ({ request }) => {
        // will create a new chat in database and return uuid
        const data = await request.formData();
        const username = data.get("name");

        let user = await fetch(
            "http://localhost:8081/users",
            { method: "POST", body: JSON.stringify({ "name": username }) }
        ).then(res => res.json());

        console.log(user)
        return { user }
    }

} satisfies Actions;
