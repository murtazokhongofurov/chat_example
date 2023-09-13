package postgres

import pbch "github.com/kafka_example/chat_service/genproto/chat_service"

func (r storagePg) AddUser(req *pbch.UserReq) (*pbch.UserRes, error) {
	res := pbch.UserRes{}
	query := `
	INSERT INTO users(first_name, last_name, user_name, bio, phone, image)
	VALUES($1, $2, $3, $4, $5, $6) RETURNING id, first_name, last_name, user_name, bio, phone, image`
	err := r.db.QueryRow(query, req.FirstName, req.LastName, req.UserName, req.Bio, req.Phone, req.Image).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.UserName, &res.Bio, &res.Phone, &res.Image)
	if err != nil {
		return &pbch.UserRes{}, err
	}
	return &res, nil
}

func (r storagePg) FindUser(req *pbch.UserId) (*pbch.UserRes, error) {
	res := pbch.UserRes{}
	query := `
	SELECT id, first_name, last_name, user_name, bio, phone, image FROM users WHERE id=$1`
	err := r.db.QueryRow(query, req.UserId).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.UserName, &res.Bio, &res.Phone, &res.Image)
	if err != nil {
		return &pbch.UserRes{}, err
	}
	return &res, nil
}

func (r storagePg) RemoveUser(req *pbch.UserId) error {
	_, err := r.db.Exec(`DELETE FROM users WHERE id=$1`, req.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r storagePg) Update(req *pbch.UserRes) (*pbch.UserRes, error) {
	res := pbch.UserRes{}
	query := `
	UPDATE users 
	SET 
		first_name=$1, last_name=$2, user_name=$3, bio=$4, phone=$5, image=$6 WHERE id=$7 
	RETURNING id, first_name, last_name, user_name, bio, phone, image`
	err := r.db.QueryRow(query, req.FirstName, req.LastName, req.UserName, req.Bio, req.Phone, req.Image, req.Id).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.UserName, &res.Bio, &res.Phone, &res.Image)
	if err != nil {
		return &pbch.UserRes{}, err
	}
	return &res, nil
}
