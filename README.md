# U-BOOT Logo Maker

Convert AI image to U-BOOT Logo format.

## STEP 1

Generate a image

![](./logo.png)

## STEP 2

Convert to U-BOOT logo format.

```bash
docker pull soulteary/onecloud-uboot:logo-maker-2025.01.06
docker run --rm -v `pwd`:/app soulteary/onecloud-uboot:logo-maker-2025.01.06 /app/logo.png
```

Or use golang run command:

```bash
go run main.go logo.png
```

## STEP 3

Get your new image.

```bash
bootup.bmp
```