import type { Meta, StoryObj } from '@storybook/svelte';

import CheckBox from './CheckBox.svelte';

const meta: Meta = {
	title: 'Form/CheckBox',
	component: CheckBox,
	argTypes: {
		label: { control: 'text' },
		checked: { control: 'boolean' }
	}
} satisfies Meta<CheckBox>;

export default meta;
type Story = StoryObj<typeof CheckBox>;

export const Primary: Story = {
	args: {
		label: 'CheckBox',
		checked: false
	}
};
