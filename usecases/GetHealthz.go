package usecases

import "context"

func (u *usecases) GetHealthz(
	ctx context.Context,
	version string,
) (map[string]string, error) {
	output := make(map[string]string)
	output["message"] = "stable"
	output["version"] = version
	return output, nil
}
