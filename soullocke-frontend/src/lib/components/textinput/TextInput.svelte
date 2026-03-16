<script lang="ts">
	import type { InputType } from '$components/textinput/types';

	type Props = {
		value: string
		label: string
		placeholder?: string
		type?: InputType
		disabled?: boolean
	};

	let {
		value = $bindable<string>(),
		label,
		placeholder,
		type = 'text',
		disabled = false,
	} : Props = $props();
</script>


<label class="label">
	<span class="label-text">{label}</span>
	<input type={type} {disabled} class="input p-3 rounded-md ring-1 hover:ring-surface-400" bind:value={value} placeholder={placeholder}/>
</label>
