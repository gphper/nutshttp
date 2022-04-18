package nutshttp

import (
	"github.com/xujiajun/nutsdb"
)

func (c *core) listSet(bucket string, key string) (data []string, err error) {

	var dataByte [][]byte

	err = c.db.View(func(tx *nutsdb.Tx) error {

		if dataByte, err = tx.SMembers(bucket, []byte(key)); err != nil {

			return err
		}

		data = make([]string, len(dataByte))
		for k, v := range dataByte {
			data[k] = string(v)
		}

		return err
	})

	return data, err
}

func (c *core) addSet(bucket string, key string, items ...string) error {
	err := c.db.Update(func(tx *nutsdb.Tx) error {

		if len(items) > 0 {
			itemsByte := make([][]byte, len(items))
			for k, v := range items {
				itemsByte[k] = []byte(v)
			}
			if err := tx.SAdd(bucket, []byte(key), itemsByte...); err != nil {

				return err
			}
		}

		return nil
	})

	return err
}

func (c *core) sAreMembers(bucket string, key string, items ...string) (ok bool, err error) {

	if err = c.db.View(
		func(tx *nutsdb.Tx) error {

			if len(items) > 0 {
				itemsByte := make([][]byte, len(items))
				for k, v := range items {
					itemsByte[k] = []byte(v)
				}

				if ok, err = tx.SAreMembers(bucket, []byte(key), itemsByte...); err != nil {
					return err
				}

			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sIsMember(bucket string, key string, item string) (ok bool, err error) {

	if err = c.db.View(

		func(tx *nutsdb.Tx) error {
			ok, err = tx.SIsMember(bucket, []byte(key), []byte(item))
			if err != nil {
				return err
			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sCard(bucket string, key string) (num int, err error) {

	if err := c.db.View(

		func(tx *nutsdb.Tx) error {

			num, err = tx.SCard(bucket, []byte(key))
			if err != nil {
				return err
			}

			return nil
		}); err != nil {

		return 0, err
	}

	return num, err
}

func (c *core) sHasKey(bucket string, key string) (ok bool, err error) {

	if err = c.db.View(

		func(tx *nutsdb.Tx) error {
			ok, err = tx.SHasKey(bucket, []byte(key))
			if err != nil {
				return err
			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sPop(bucket string, key string) (item string, err error) {

	if err = c.db.Update(
		func(tx *nutsdb.Tx) error {

			itemByte, err := tx.SPop(bucket, []byte(key))
			if err != nil {
				return err
			}
			item = string(itemByte)

			return nil

		}); err != nil {

		return
	}

	return
}

func (c *core) sRem(bucket string, key string, items []string) error {

	if err := c.db.Update(

		func(tx *nutsdb.Tx) error {

			if len(items) > 0 {

				itemsByte := make([][]byte, len(items))
				for k, v := range items {
					itemsByte[k] = []byte(v)
				}

				if err := tx.SRem(bucket, []byte(key), itemsByte...); err != nil {
					return err
				}
			}

			return nil
		}); err != nil {

		return err
	}

	return nil
}

func (c *core) sDiffByOneBucket(bucket string, key1 string, key2 string) (items []string, err error) {

	if err = c.db.View(

		func(tx *nutsdb.Tx) error {
			itemsBytes, err := tx.SDiffByOneBucket(bucket, []byte(key1), []byte(key2))
			if err != nil {

				return err
			}

			if len(itemsBytes) > 0 {
				items = make([]string, len(itemsBytes))
				for k, itemByte := range itemsBytes {
					items[k] = string(itemByte)
				}
			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sDiffByTwoBuckets(bucket1 string, bucket2 string, key1 string, key2 string) (items []string, err error) {

	if err = c.db.View(

		func(tx *nutsdb.Tx) error {
			itemsBytes, err := tx.SDiffByTwoBuckets(bucket1, []byte(key1), bucket2, []byte(key2))
			if err != nil {

				return err
			}

			if len(itemsBytes) > 0 {
				items = make([]string, len(itemsBytes))
				for k, itemByte := range itemsBytes {
					items[k] = string(itemByte)
				}
			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sMoveByOneBucket(bucket string, key1 string, key2 string, item string) (ok bool, err error) {

	if err = c.db.Update(
		func(tx *nutsdb.Tx) error {
			ok, err = tx.SMoveByOneBucket(bucket, []byte(key1), []byte(key2), []byte(item))
			if err != nil {

				return err
			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sMoveByTwoBuckets(bucket string, key1 string, bucket2 string, key2 string, item string) (ok bool, err error) {

	if err = c.db.Update(
		func(tx *nutsdb.Tx) error {
			ok, err = tx.SMoveByTwoBuckets(bucket, []byte(key1), bucket2, []byte(key2), []byte(item))
			if err != nil {

				return err
			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sUnionByOneBucket(bucket string, key1 string, key2 string) (items []string, err error) {

	if err = c.db.View(

		func(tx *nutsdb.Tx) error {
			itemBytes, err := tx.SUnionByOneBucket(bucket, []byte(key1), []byte(key2))
			if err != nil {

				return err
			}

			if len(itemBytes) > 0 {
				items = make([]string, len(itemBytes))
				for k, itemByte := range itemBytes {
					items[k] = string(itemByte)
				}
			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sUnionByTwoBuckets(bucket1 string, key1 string, bucket2 string, key2 string) (items []string, err error) {

	if err = c.db.View(

		func(tx *nutsdb.Tx) error {
			itemBytes, err := tx.SUnionByTwoBuckets(bucket1, []byte(key1), bucket2, []byte(key2))
			if err != nil {

				return err
			}

			if len(itemBytes) > 0 {
				items = make([]string, len(itemBytes))
				for k, itemByte := range itemBytes {
					items[k] = string(itemByte)
				}
			}

			return nil
		}); err != nil {

		return
	}

	return
}
