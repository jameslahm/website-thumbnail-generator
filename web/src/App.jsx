import React, { useState, useEffect } from 'react';
import './App.css';
import {
  Box,
  Button,
  ChakraProvider,
  Heading,
  Input,
  Image,
} from '@chakra-ui/react';
import theme from './theme';

function App() {
  const [url, setUrl] = useState('');
  const [screenshot, setScreenShot] = useState('');

  const handleSubmit = () => {
    if (url) {
      fetch('http://127.0.0.1:3000/api/thumbnail', {
        headers: {
          'Content-Type': 'application/json',
        },
        method: 'POST',
        body: JSON.stringify({ url }),
      })
        .then((res) => res.json())
        .then((data) => {
          setScreenShot(data.screenshot);
        });
    }
  };

  return (
    <ChakraProvider theme={theme}>
      <Box
        h={'100%'}
        display="flex"
        justifyContent="space-around"
        alignItems="center"
        flexDirection="column"
      >
        <Box>
          <Heading>Generate a thumbnail of a website</Heading>
          <Input
            maxW="xl"
            mt={4}
            value={url}
            placeholder="Your Url"
            onChange={(e) => setUrl(e.target.value)}
          ></Input>
          <Box w={'100%'}>
            <Button
              onClick={handleSubmit}
              mt={2}
              colorScheme="teal"
              variant="solid"
            >
              Generate
            </Button>
          </Box>
          {screenshot ? <Image mt={3} src={screenshot}></Image> : null}
        </Box>
      </Box>
    </ChakraProvider>
  );
}

export default App;
