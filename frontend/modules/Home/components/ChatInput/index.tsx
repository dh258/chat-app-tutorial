import { ChangeEventHandler } from 'react';
import { Input } from '@mui/material';
import Container from './Container';

type Props = {
  value: string;
  onChange: ChangeEventHandler<HTMLTextAreaElement | HTMLInputElement>;
};

const ChatInput = ({ value, onChange }: Props) => {
  return (
    <Container>
      <Input value={value} onChange={onChange} />
    </Container>
  );
};

export default ChatInput;
