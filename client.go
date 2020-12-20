package quantd

import (
	"image"
	"os/exec"

	quantdimage "github.com/1lann/quantd/image"
	"github.com/tinylib/msgp/msgp"
)

type Client struct {
	cmd    *exec.Cmd
	input  *msgp.Writer
	output *msgp.Reader
}

func (e *Executable) NewClient() (*Client, error) {
	cmd := exec.Command(e.path, "start")

	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	inPipe, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	out := msgp.NewReader(outPipe)
	in := msgp.NewWriter(inPipe)

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return &Client{
		cmd:    cmd,
		input:  in,
		output: out,
	}, nil
}

func (c *Client) Quantize(img image.Image, speed, maxColors int, dither float64) (*quantdimage.ImageResponse, error) {
	req := quantdimage.NewImageRequest(img, speed, maxColors, dither)

	err := req.EncodeMsg(c.input)
	if err != nil {
		return nil, err
	}
	err = c.input.Flush()
	if err != nil {
		return nil, err
	}

	var resp quantdimage.ImageResponse
	err = resp.DecodeMsg(c.output)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) Stop() {
	c.cmd.Process.Kill()
	c.cmd.Process.Wait()
}
