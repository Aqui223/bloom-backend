package KeysRepo

func (k *KeysRepo) Delete(id int) error {
	query := `DELETE FROM keys WHERE id = $1`

	_, err := k.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
