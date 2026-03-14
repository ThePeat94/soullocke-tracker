<script lang="ts">

	import type { VariantType } from './types.ts';
	import type { PropsWithChildren } from '$util/types';

	const variantClasses: Record<VariantType, string> = {
		filled: 'preset-filled-primary-500',
		outlined: 'preset-outlined-primary-500',
	};

	type Props = {
		variant : VariantType,
		onClick : () => void,
	};

	let {
		variant,
		onClick,
		children,
	} : PropsWithChildren<Props> = $props()

</script>

<button type="button" class={`btn ${variantClasses[variant]}`} onclick={onClick}>{@render children()}</button>
