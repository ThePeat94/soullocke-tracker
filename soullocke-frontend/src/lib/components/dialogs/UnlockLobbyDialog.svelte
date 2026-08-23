<script lang="ts">

	import Card from '$components/Card.svelte';
	import PrimaryButton from '$components/buttons/PrimaryButton.svelte';
	import TextInput from '$components/textinput/TextInput.svelte';
	import { getAuthMutation } from '$api/auth';

	let { lobbyId } = $props();

	let password = $state('');
	let dialog = $state<HTMLDialogElement>();
	const loginMutation = getAuthMutation();


	const handleUnlockLobbyClick = async (): Promise<void> => {
			await loginMutation.mutateAsync({
				body: {
					lobbyId,
					password,
				},
			});
			handleDialogClose();
	};

	export const openDialog = (): void => {
		dialog?.showModal();
	};

	const handleDialogClose = (): void => {
		password = '';
		dialog?.close();
	};

</script>

<dialog
	bind:this={dialog}
	onclose={handleDialogClose}
	class="m-auto"
>
	<Card>
		{#snippet header()}
			<h2 class="h4">Unlock lobby</h2>
		{/snippet}
		{#snippet content()}
			<div class="flex flex-col gap-4">
				<TextInput
					label="Password"
					type="password"
					bind:value={password}
				/>
			</div>
		{/snippet}
		{#snippet footer()}
			<PrimaryButton onClick={handleUnlockLobbyClick} variant="filled" disabled={loginMutation.isPending} >Join Lobby</PrimaryButton>
			<PrimaryButton onClick={handleDialogClose} variant="filled">Cancel</PrimaryButton>
		{/snippet}
	</Card>
</dialog>
