<script lang="ts">
	import { enhance } from '$app/forms';
	import Button from '$lib/components/Button/Button.svelte';
	import { onMount } from 'svelte';
	import type { ActionData } from '../../$types';

	export let form: ActionData;

	let mounted: boolean = false;
	let username: string | null = '';
	onMount(async () => {
		mounted = true;
		username = localStorage.getItem('userName');
		if (username) {
			window.location.href = '/account';
		}
	});

	async function linkUser(chatID: string, userID: string) {
		console.log('linking user');
		const link = await fetch(`http://localhost:8081/chats/link`, {
			method: 'POST',
			body: JSON.stringify({ chat_id: chatID, user_id: userID })
		}).then((res) => {
			res.status, res.statusText;
		});
		console.log(link);
		// redirect to chatpage
		window.location.href = `/chat/${chatID}`;
		localStorage.removeItem('chatId');
	}

	$: {
		if (form?.user && mounted) {
			localStorage.setItem('userName', form.user.name);
			localStorage.setItem('userID', form.user.user_id);
			window.location.href = '/';
			username = form.user.username;
			const chatID = localStorage.getItem('chatId');
			if (chatID) {
				linkUser(chatID, form.user.user_id);
			}
		}
	}
</script>

<p>No User Identifier has been found in your browser, Please Create one below</p>
<div class="flex flex-col gap-1.5 pt-4">
	<p>Create one below</p>
	<form action="?/createUser" method="POST" use:enhance class="flex flex-row gap-2 items-center">
		<label for="name">User Identifier</label>
		<input
			class="p-1.5"
			type="text"
			name="name"
			id="user"
			placeholder="Enter a user identifier"
			required
		/>
		<button type="submit" on:click={() => {}}><Button size="small">Create Account</Button></button>
	</form>
</div>
