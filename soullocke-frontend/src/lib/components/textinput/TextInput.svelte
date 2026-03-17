<script lang="ts">
	import type { InputType } from '$components/textinput/types';

	type Props = {
		value: string
		label: string
		placeholder?: string
		type?: InputType
		disabled?: boolean
		minLength?: number,
		maxLength?: number,
		valid?: boolean,
	};

	let {
		value = $bindable<string>(),
		label,
		placeholder,
		type = 'text',
		disabled = false,
		minLength,
		maxLength,
		valid = $bindable<boolean>(true)
	}: Props = $props();

	let dirty = $state(false);

	const tooLong = $derived(maxLength && value.length > maxLength);
	const tooShort = $derived(minLength && value.length < minLength);
	$effect(() => {
		valid = !tooLong && !tooShort;
	});

	const getErrorMessage = $derived(() => {
		if (tooShort) {
			return `Must be at least ${minLength} characters.`;
		}
		if (tooLong) {
			return `Must be at most ${maxLength} characters.`;
		}
		return undefined;
	});
</script>

<label class="label">
	<span class="label-text" class:text-error-500={!valid && dirty}>{label}</span>
	<div class="relative">
		<input
			type={type}
			{disabled}
			class="input p-3 pr-20 rounded-md ring-1 hover:ring-surface-400"
			{placeholder}
			class:ring-red-500={(tooLong || tooShort) && dirty}
			class:valid={valid}
			minlength={minLength}
			oninput={() => dirty = true}
			bind:value={value}
		/>
		<span
			class="absolute right-3 top-1/2 -translate-y-1/2 text-sm"
			class:text-success-500={valid}
			class:text-error-500={!valid && dirty}
		>
			{value.length}{#if maxLength}/{maxLength}{/if}
		</span>
	</div>
	{#if dirty && !valid}
		<span class="text-error-500 text-xs">{getErrorMessage()}</span>
	{/if}
</label>
