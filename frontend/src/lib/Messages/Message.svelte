<script lang="ts">
	import { users, type message } from '$lib/Messages/messagesStores';
	import { timeDisplay } from '$lib/utils/timeDisplay';

	export let message: message;
	import MessageSettings from './MessageSettings.svelte';
	import Reactions from './Reactions.svelte';
</script>

<div class="flex flex-col bg-red-50 p-2 m-4">
	<div class="bg-blue-50">
		<div class="font-semibold text-lg">
			{$users[message.user_id]}
		</div>
		{#if !message.deleted}
			<div class="content">
				{message.body}
			</div>
		{:else}
			<div class="content">Message has been deleted</div>
		{/if}
	</div>
	<div class="flex flex-row justify-between bg-blue-100 gap-2">
		<div class="">
			<Reactions {message} />
		</div>
		<div class="flex flex-row justify-end gap-2">
			{timeDisplay(message.created_at)}
			{#if message.updated_at != message.created_at}
				{#if message.deleted}
					(deleted {timeDisplay(message.updated_at)})
				{:else}
					(edited {timeDisplay(message.updated_at)})
				{/if}
			{/if}
			<MessageSettings {message} />
		</div>
	</div>
</div>
